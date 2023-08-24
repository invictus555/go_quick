package service

type ReferenceType int32

const (
	IsAdmin      ReferenceType = 1   // direct reference type
	HasAdmin     ReferenceType = -1  // direct reference type
	IsChild      ReferenceType = 2   // direct reference type
	IsParent     ReferenceType = -2  // direct reference type
	IsDescendant ReferenceType = 20  // indirect
	IsAncestor   ReferenceType = -20 // indirect
	IsReport     ReferenceType = 3   // direct reference type
	IsLeader     ReferenceType = -3  // direct reference type
)

type OperationType int32

const (
	Create OperationType = 0
	Read   OperationType = 1
	Update OperationType = 2
	Delete OperationType = 3
)

type EntityType int32

const (
	IdentityPrincipal EntityType = 0
	RoleAction        EntityType = 1
	ResourceObject    EntityType = 2
)

type HierarchyRelation int64

const (
	Subset     HierarchyRelation = 0
	Superset   HierarchyRelation = 1
	Equivalent HierarchyRelation = 2
	Overlap    HierarchyRelation = 3
)

type Pair struct {
	Key string
	Val string
}
type PolicyEntity struct {
	ID         int64                    `json:"id,omitempty" validate:"omitempty,gte=1"`
	Version    int64                    `json:"version,omitempty" validate:"omitempty,gte=1"`
	Ns         string                   `json:"ns,omitempty" validate:"omitempty,lte=64,ns"`
	Location   string                   `json:"location,omitempty" validate:"omitempty,lte=64,location"`
	Path       string                   `json:"path,omitempty" validate:"omitempty,lte=256,path"`
	Metadata   []map[string]interface{} `json:"metadata,omitempty" validate:"omitempty,maxSize=1024"`
	Attributes map[string]interface{}   `json:"attributes,omitempty" validate:"omitempty,maxSize=8192"`
	PathV2     []*Pair                  `json:"pathV2,omitempty" validate:"omitempty,maxSize=1024"`
	LocationV2 []*Pair                  `json:"locationV2,omitempty" validate:"omitempty,maxSize=1024"`
	IsEnable   bool                     `json:"is_enable,omitempty"`
	CreateTime int64                    `json:"create_time,omitempty" validate:"omitempty,gte=1"`
}

type Policy struct {
	ID           int64                    `json:"id,omitempty" validate:"omitempty,gte=1"`
	Version      int64                    `json:"version,omitempty" validate:"omitempty,gte=1"`
	Identity     *PolicyEntity            `json:"identity,omitempty" validate:"omitempty,dive"`
	IdentityType int64                    `json:"identity_type"`
	Role         *PolicyEntity            `json:"role,omitempty" validate:"omitempty,dive"`
	RoleType     int64                    `json:"role_type,omitempty"`
	Resource     *PolicyEntity            `json:"resource,omitempty" validate:"omitempty,dive"`
	ResourceType int64                    `json:"resource_type,omitempty"`
	ExpireTime   int64                    `json:"expire_time"`
	Condition    []map[string]interface{} `json:"condition,omitempty" validate:"omitempty,maxSize=2048"`
	IsEnable     bool                     `json:"is_enable,omitempty"`
	IsBreakGlass bool                     `json:"is_break_glass,omitempty"`
	Information  map[string]interface{}   `json:"information,omitempty" validate:"omitempty"`
	CreateTime   int64                    `json:"create_time,omitempty" validate:"omitempty,gte=1"`
}

type Query struct {
	Operation            OperationType                `json:"operation"`
	EntityType           EntityType                   `json:"entity_type"`
	Entities             []*PolicyEntity              `json:"entities,omitempty"`
	BaseEntity           *PolicyEntity                `json:"base_entity,omitempty"`
	Policies             []*Policy                    `json:"policies,omitempty"`
	BasePolicy           *Policy                      `json:"base_policy,omitempty"`
	EntitiesWithRefs     []*PolicyEntityWithReference `json:"entities_with_refs,omitempty"`
	StartID              int64                        `json:"start_id,omitempty"`
	BatchSize            int64                        `json:"batch_size,omitempty"`
	Relation             HierarchyRelation            `json:"relation,omitempty"`
	UpdateOnExist        bool                         `json:"update_on_exist,omitempty"`
	ReturnCreatedResults bool                         `json:"return_created_results,omitempty"`
	IgnoreNotFound       bool                         `json:"ignore_not_found,omitempty"`
	HardDelete           bool                         `json:"hard_delete,omitempty"`
	AsyncFixNestingRefs  bool                         `json:"async_fix_nesting_refs,omitempty"`
	UpsertNestingRefs    bool                         `json:"upsert_nesting_refs,omitempty"`
}

type RefInfo struct {
	ExpireTime  int64                    `json:"expire_time"`
	IsEnable    bool                     `json:"is_enable,omitempty"`
	Condition   []map[string]interface{} `json:"condition,omitempty" validate:"omitempty,maxSize=2048"`
	Information map[string]interface{}   `json:"information,omitempty" validate:"omitempty"`
	CreateTime  int64                    `json:"create_time,omitempty" validate:"omitempty,gte=1"`
	Version     int64                    `json:"version,omitempty" validate:"omitempty,gte=1"`
} // 这里的fields都是和policy的是一样的

type ReferenceEntity struct {
	Entity      *PolicyEntity `json:"entity,omitempty"`
	EntityType  EntityType    `json:"entity_type"`
	RefInfo     *RefInfo      `json:"ref_info,omitempty"`
	BaseRefInfo *RefInfo      `json:"base_ref_info,omitempty"`
}

type Reference struct {
	Operation            OperationType      `json:"operation"`
	RefType              ReferenceType      `json:"ref_type,omitempty"`
	BaseRefType          ReferenceType      `json:"base_ref_type,omitempty"`
	RefEntities          []*ReferenceEntity `json:"ref_entities,omitempty"`
	StartID              int64              `json:"start_id,omitempty"`
	BatchSize            int64              `json:"batch_size,omitempty"`
	UpdateOnExist        bool               `json:"update_on_exist,omitempty"`
	ReturnCreatedResults bool               `json:"return_created_results,omitempty"`
	IgnoreNotFound       bool               `json:"ignore_not_found,omitempty"`
	HardDelete           bool               `json:"hard_delete,omitempty"`
}

type PolicyEntityWithReference struct {
	Entity     *PolicyEntity `json:"entity,omitempty"`
	BaseEntity *PolicyEntity `json:"base_entity,omitempty"`
	References []*Reference  `json:"references,omitempty"`
}

type OperateTransactionRequest struct {
	Queries []*Query `json:"queries,omitempty"`
}

type PolicyOperation struct {
	Policies             []*Policy         `json:"policies,omitempty" validate:"lte=1000000,dive"`
	BasePolicy           *Policy           `json:"base_policy,omitempty"`
	StartID              int64             `json:"start_id,omitempty"`
	BatchSize            int64             `json:"batch_size,omitempty" validate:"omitempty,lte=500000"`
	Relation             HierarchyRelation `json:"relation"`
	UpdateOnExist        bool              `json:"update_on_exist,omitempty"`
	ReturnCreatedResults bool              `json:"return_created_results,omitempty"`
	IgnoreNotFound       bool              `json:"ignore_not_found,omitempty"`
}

type GetPermissionOption struct {
	IsStrongConsistency          bool `json:"is_strong_consistency"`
	IsExactMatch                 bool `json:"is_exact_match"`
	IsFromGetAllowedPolicyEntity bool `json:"is_from_get_allowed_policy_entity,omitempty"`
}

type TaskEntity struct {
	Ns         string                   `json:"ns" validate:"omitempty,lte=64,ns"`
	Location   string                   `json:"location,omitempty" validate:"omitempty,lte=64,location"`
	Path       string                   `json:"path,omitempty" validate:"omitempty,lte=256,path"`
	Metadata   []map[string]interface{} `json:"metadata,omitempty" validate:"omitempty,maxSize=1024"`
	Attributes map[string]interface{}   `json:"attributes,omitempty" validate:"omitempty,maxSize=8192"`
	PathV2     []*Pair                  `form:"pathV2" json:"pathV2" validate:"omitempty,maxSize=1024"`
	LocationV2 []*Pair                  `form:"locationV2" json:"locationV2" validate:"omitempty,maxSize=1024"`
}

type TaskConfig struct {
	StartID                     int64    `json:"start_id,omitempty"`
	BatchSize                   int64    `json:"batch_size,omitempty" validate:"omitempty,lte=500000"`
	SkipIdentityOverlapFetch    bool     `json:"skip_identity_overlap_fetch,omitempty"`
	SkipRoleOverlapFetch        bool     `json:"skip_role_overlap_fetch,omitempty"`
	SkipResourceOverlapFetch    bool     `json:"skip_resource_overlap_fetch,omitempty"`
	IdentityReferenceNamespaces []string `json:"identity_reference_namespaces,omitempty" validate:"omitempty,dive,lte=64,ns"`
	RoleReferenceNamespaces     []string `json:"role_reference_namespaces,omitempty" validate:"omitempty,dive,lte=64,ns"`
	ResourceReferenceNamespaces []string `json:"resource_reference_namespaces,omitempty" validate:"omitempty,dive,lte=64,ns"`
	SkipIdentityLgroupFetch     bool     `json:"skip_identity_lgroup_fetch,omitempty"`
	SkipRoleLgroupFetch         bool     `json:"skip_role_lgroup_fetch,omitempty"`
	SkipResourceLgroupFetch     bool     `json:"skip_resource_lgroup_fetch,omitempty"`
	IdentityExactMatch          bool     `json:"identity_exact_match,omitempty"`
	RoleExactMatch              bool     `json:"role_exact_match,omitempty"`
	ResourceExactMatch          bool     `json:"resource_exact_match,omitempty"`
	IdentityNeedDisabled        bool     `json:"identity_need_disabled,omitempty"`
	RoleNeedDisabled            bool     `json:"role_need_disabled,omitempty"`
	ResourceNeedDisabled        bool     `json:"resource_need_disabled,omitempty"`
	PolicyNeedDisabled          bool     `json:"policy_need_disabled,omitempty"`
	IdentityReferenceStartId    int64    `json:"identity_reference_start_id,omitempty"`
	IdentityReferenceBatchSize  int64    `json:"identity_reference_batch_size,omitempty"`
	RoleReferenceStartId        int64    `json:"role_reference_start_id,omitempty"`
	RoleReferenceBatchSize      int64    `json:"role_reference_batch_size,omitempty"`
	ResourceReferenceStartId    int64    `json:"resource_reference_start_id,omitempty"`
	ResourceReferenceBatchSize  int64    `json:"resource_reference_batch_size,omitempty"`
}

type Task struct { // 鉴权
	AuthenticationID string      `json:"authentication_id,omitempty" validate:"omitempty,spiffe"`
	Action           *TaskEntity `json:"action,omitempty" validate:"omitempty,dive"`
	Object           *TaskEntity `json:"object,omitempty" validate:"omitempty,dive"`
	Principal        *TaskEntity `json:"principal,omitempty" validate:"omitempty,dive"`
	Config           *TaskConfig `json:"config,omitempty" validate:"omitempty,dive"`
}

type RequestContext struct {
	Timestamp   int64  `json:"timestamp,omitempty" validate:"omitempty,gte=0"`
	PrincipalIP string `json:"principal_ip,omitempty" validate:"omitempty,ip"`
}

type RequestConfig struct {
	NeedDisabled bool `json:"need_disabled"`
	IsAudit      bool `json:"is_audit"` // IsAudit is use to decide whether we should log the request as a permission check in SDK
}
type GetPermissionReq struct {
	Task    *Task           `json:"task,omitempty" validate:"omitempty,dive"`
	Context *RequestContext `json:"context,omitempty" validate:"omitempty,dive"`
	Config  *RequestConfig  `json:"config,omitempty" validate:"omitempty,dive"`
}

type GetBatchPermissionsReq struct {
	Tasks   []*Task         `json:"tasks" validate:"omitempty,dive"`
	Context *RequestContext `json:"context,omitempty" validate:"omitempty,dive"`
	Config  *RequestConfig  `json:"config,omitempty" validate:"omitempty,dive"`
}

type GetPermissionRequest struct {
	Request *GetPermissionReq    `json:"request" validate:"required,dive"`
	Option  *GetPermissionOption `json:"option,omitempty" validate:"omitempty,dive"`
}

type GetBatchPermissionsRequest struct {
	Request *GetBatchPermissionsReq `json:"request" validate:"required,dive"`
	Option  *GetPermissionOption    `json:"option,omitempty" validate:"omitempty,dive"`
}
