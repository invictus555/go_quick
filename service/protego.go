package service

import (
	"encoding/json"
	"fmt"
	"quick_practices/utils"
	"time"

	"github.com/pkg/errors"
)

const (
	ProtegoURL    = "https://protego-api.byted.org"
	ProtegoBoeURL = "https://protego-boe-api.byted.org/"
)

// TODO  tcc配置
const (
	AppID     = ""
	AppSecret = ""
)

// RegisterResource 注册资源、添加添加管理员
func RegisterResource(resourceKey, namespace, region string, admins, members []string) error {
	regionSettings := []*Pair{
		&Pair{
			Key: "r",
			Val: region,
		},
	}
	if region == "boe" {
		regionSettings = append(regionSettings, &Pair{
			Key: "vdc",
			Val: "boe",
		})
	}

	// 定义资源
	resourceQuery := &Query{
		EntityType:           ResourceObject,
		Operation:            Create,
		UpdateOnExist:        true,
		ReturnCreatedResults: true,
		EntitiesWithRefs: []*PolicyEntityWithReference{
			&PolicyEntityWithReference{
				Entity: &PolicyEntity{
					IsEnable:   true,
					Ns:         namespace,
					LocationV2: regionSettings,
					PathV2: []*Pair{
						&Pair{
							Key: "key",
							Val: resourceKey,
						},
					},
					Attributes: map[string]interface{}{
						"security_level":  3, // 安全等级"
						"auth_block_type": 0,
						"product": map[string]interface{}{
							"default": "product_toutiaoyanfa-shipinjiagou",
						},
						"audit_status":      0,
						"audit_periodicity": 360,
						"authorizable":      1,
						"region_type":       0,
						"name":              resourceKey, // 资源名称
					},
				},
				References: getAdminReferences(admins), // 只关联管理员
			},
		},
	}

	// 定义动作属性--资源有两种属性(Read and Read-write)
	attributesQuery := &Query{
		EntityType:           RoleAction,
		Operation:            Create,
		Relation:             Equivalent,
		UpdateOnExist:        true,
		ReturnCreatedResults: true,
		Entities: []*PolicyEntity{
			&PolicyEntity{
				IsEnable:   true,
				Ns:         namespace,
				LocationV2: regionSettings,
				PathV2: []*Pair{
					&Pair{
						Key: "resource",
						Val: resourceKey,
					},
					&Pair{
						Key: "action",
						Val: "Read_Write",
					},
				},
				Attributes: map[string]interface{}{
					"name": "读写权限",
				},
			},
			&PolicyEntity{
				IsEnable:   true,
				Ns:         namespace,
				LocationV2: regionSettings,
				PathV2: []*Pair{
					&Pair{
						Key: "resource",
						Val: resourceKey,
					},
					&Pair{
						Key: "action",
						Val: "Read",
					},
				},
				Attributes: map[string]interface{}{
					"name": "读权限",
				},
			},
		},
	}

	// 构建事务级请求
	req := &OperateTransactionRequest{
		Queries: []*Query{
			resourceQuery,
			attributesQuery,
		},
	}

	body, _ := json.Marshal(req)
	fmt.Println(string(body))
	msg, err := utils.DoHttpPostWithBasicAuth(ProtegoURL+"/transaction", AppID, AppSecret, body) // 事务级操纵注册资源，添加管理员/成员并赋予对应的权限
	if err != nil || msg == nil {
		fmt.Println("failed:", string(msg), err)
		return errors.New(string(msg))
	}
	fmt.Println("resp:", string(msg))
	return nil
}

// AddAdmins 在注册资源之后新增管理员
func AddAdmins(namespace, region, resourceKey string, reqs []*AuthorizationRequest) error {
	if reqs == nil || len(namespace) == 0 || len(resourceKey) == 0 {
		return utils.KaniInputParametersError
	}

	for _, req := range reqs {
		if req == nil {
			continue
		}

		req.Attributes = append(req.Attributes, "admin")
	}

	return AuthorizeBatch(namespace, region, resourceKey, reqs)
}

// Authorize [单个] 给管理员授权/新增非管理员用户并授权
func Authorize(user, namespace, resourceKey, region string, duration time.Duration, attributes []string) error {
	reqs := []*AuthorizationRequest{
		&AuthorizationRequest{
			User:       user,
			Duration:   duration,
			Attributes: attributes,
		},
	}

	return AuthorizeBatch(namespace, region, resourceKey, reqs)
}

// AuthorizeBatch [批量] 给管理员授权或者新增非管理员并给予授权
func AuthorizeBatch(namespace, region, resourceKey string, reqs []*AuthorizationRequest) error {
	if reqs == nil || len(namespace) == 0 || len(region) == 0 || len(reqs) == 0 {
		return errors.New("input parameters are illegal")
	}

	regionSettings := []*Pair{
		&Pair{
			Key: "r",
			Val: region,
		},
	}
	if region == "boe" {
		regionSettings = append(regionSettings, &Pair{
			Key: "vdc",
			Val: "boe",
		})
	}

	policyOperation := &PolicyOperation{
		UpdateOnExist: true, // 用于强制更新
		Relation:      Equivalent,
		Policies:      getPolicies(namespace, resourceKey, regionSettings, reqs),
	}

	body, _ := json.Marshal(policyOperation)
	fmt.Println(string(body))
	return nil
}

// CheckPermission 单个鉴权
func CheckPermission(user, namespace, resourceKey, region, attribute string) error {
	if len(user) == 0 || len(namespace) == 0 || len(resourceKey) == 0 || len(region) == 0 || len(attribute) == 0 {
		return errors.New("CheckPermission: input parameters are illegal")
	}

	requests := []*PermissionRequest{
		&PermissionRequest{
			User:       user,
			Attributes: []string{attribute},
		},
	}

	tasks := getPermissionTasks(namespace, resourceKey, region, requests)
	req := &GetPermissionRequest{
		Option: &GetPermissionOption{
			IsExactMatch:        true,
			IsStrongConsistency: true,
		},
		Request: &GetPermissionReq{
			Task: tasks[0],
		},
	}

	body, _ := json.Marshal(req)
	fmt.Println(string(body))
	return nil
}

// CheckBatchPermissions 批量鉴权
func CheckBatchPermissions(namespace, resourceKey, region string, requests []*PermissionRequest) error {
	if len(namespace) == 0 || len(resourceKey) == 0 || len(region) == 0 || requests == nil || len(requests) == 0 {
		return errors.New("BatchCheckPermissions: input parameters are illegal")
	}

	tasks := getPermissionTasks(namespace, resourceKey, region, requests)
	if tasks == nil {
		return nil
	}

	req := &GetBatchPermissionsRequest{
		Option: &GetPermissionOption{
			IsExactMatch:        true,
			IsStrongConsistency: true,
		},
		Request: &GetBatchPermissionsReq{
			Tasks: tasks,
		},
	}

	body, _ := json.Marshal(req)
	fmt.Println(string(body))
	return nil
}
