package service

import "time"

// AuthorizationRequest 授权请求
type AuthorizationRequest struct {
	User       string        // 用户名
	Attributes []string      // 希望授予的权限
	Duration   time.Duration // 有效期
}

// PermissionRequest 鉴权请求
type PermissionRequest struct {
	User       string   // 用户名
	Attributes []string // 被鉴权的权限
}

func getAdminReferences(admins []string) []*Reference {
	var references []*Reference
	adminRefEntities := getAdminRefEntities(admins)
	if len(adminRefEntities) > 0 {
		references = append(references, &Reference{
			RefType:              HasAdmin,
			Operation:            Create,
			UpdateOnExist:        true,
			ReturnCreatedResults: true,
			RefEntities:          adminRefEntities,
		})
	}
	return references
}

func getAdminRefEntities(admins []string) []*ReferenceEntity {
	if len(admins) == 0 {
		return nil
	}

	var adminRefEntities []*ReferenceEntity
	for _, admin := range admins {
		adminRefEntities = append(adminRefEntities, &ReferenceEntity{
			EntityType: IdentityPrincipal,
			Entity: &PolicyEntity{
				Ns: "user",
				PathV2: []*Pair{
					&Pair{
						Key: "id",
						Val: admin,
					},
				},
			},
			RefInfo: &RefInfo{
				ExpireTime: 0,
				IsEnable:   true,
			},
		})
	}

	return adminRefEntities
}

func getMemberRefEntities(members []string) []*ReferenceEntity {
	var memberRefEntities []*ReferenceEntity
	for _, member := range members {
		memberRefEntities = append(memberRefEntities, &ReferenceEntity{
			EntityType: IdentityPrincipal,
			Entity: &PolicyEntity{
				Ns: "user",
				PathV2: []*Pair{
					&Pair{
						Key: "id",
						Val: member,
					},
				},
			},
			RefInfo: &RefInfo{
				ExpireTime: 0,
				IsEnable:   true,
			},
		})
	}
	return memberRefEntities
}

func getPolicies(namespace, resourceKey string, regionSettings []*Pair, reqs []*AuthorizationRequest) []*Policy {
	var policies []*Policy
	if reqs == nil || regionSettings == nil || len(namespace) == 0 || len(resourceKey) == 0 {
		return policies
	}

	for _, req := range reqs { // 多用户批量授权
		if reqs == nil || len(req.Attributes) == 0 {
			continue
		}

		for _, attr := range req.Attributes { // 构造policy
			exp := int64(time.Now().Add(req.Duration).UnixNano())
			if attr == "admin" { // 添加管理员
				policies = append(policies, &Policy{
					ExpireTime:   exp,
					IdentityType: 0,
					IsEnable:     true,
					Identity: &PolicyEntity{
						Ns: "user", //用户，必须user
						PathV2: []*Pair{
							&Pair{
								Key: "id", // 用户，必须是id
								Val: req.User,
							},
						},
					},
					RoleType: 1,
					Role: &PolicyEntity{
						Ns: "reftype_" + namespace,
						PathV2: []*Pair{
							&Pair{
								Key: "action",
								Val: "admin",
							},
						},
						LocationV2: regionSettings,
					},
					ResourceType: 2,
					Resource: &PolicyEntity{
						Ns: namespace,
						PathV2: []*Pair{
							&Pair{
								Key: "key",
								Val: resourceKey,
							},
						},
						LocationV2: regionSettings,
					},
				})
				continue
			}
			policies = append(policies, &Policy{
				ExpireTime:   exp,
				IdentityType: 0,
				Identity: &PolicyEntity{
					Ns: "user", //用户，必须user
					PathV2: []*Pair{
						&Pair{
							Key: "id", // 用户，必须是id
							Val: req.User,
						},
					},
				},
				RoleType: 1,
				Role: &PolicyEntity{
					Ns: namespace,
					PathV2: []*Pair{
						&Pair{
							Key: "resource",
							Val: resourceKey,
						},
						&Pair{
							Key: "action",
							Val: attr,
						},
					},
					LocationV2: regionSettings,
				},
				ResourceType: 2,
				Resource: &PolicyEntity{
					Ns: namespace,
					PathV2: []*Pair{
						&Pair{
							Key: "key",
							Val: resourceKey,
						},
					},
					LocationV2: regionSettings,
				},
			})
		}
	}

	return policies
}

func getPermissionTasks(namespace, resourceKey, region string, requests []*PermissionRequest) []*Task {
	var tasks []*Task
	if requests == nil || len(requests) == 0 {
		return tasks
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

	for _, request := range requests {
		if request == nil {
			continue
		}

		for _, attribute := range request.Attributes {
			tasks = append(tasks, &Task{
				Config: &TaskConfig{},
				Action: &TaskEntity{
					Ns: namespace,
					PathV2: []*Pair{ // 数组内的pair位置不能改动
						&Pair{
							Key: "resource",
							Val: resourceKey,
						},
						&Pair{
							Key: "action",
							Val: attribute,
						},
					},
					LocationV2: regionSettings,
				},
				Object: &TaskEntity{
					Ns: namespace,
					PathV2: []*Pair{
						&Pair{
							Key: "key",
							Val: resourceKey,
						},
					},
					LocationV2: regionSettings,
				},
				Principal: &TaskEntity{
					Ns: "user",
					PathV2: []*Pair{
						&Pair{
							Key: "id",
							Val: request.User,
						},
					},
					LocationV2: regionSettings,
				},
			})
		}
	}

	return tasks
}
