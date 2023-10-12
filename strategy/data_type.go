package strategy

var RuleStageID int64 = 1

type RuleGroupInfo struct {
	ID           int64              `json:"id"`            // rule group id
	TenantID     int64              `json:"tenant_id"`     // current tenant id
	GroupName    string             `json:"group_name"`    // page card name
	Description  string             `json:"description"`   // page description
	RuleStageID  *int64             `json:"rule_stage_id"` // rule stage id
	CountryGroup []CountryGroupInfo `json:"country_group"` // country group
}

type CountryGroupInfo struct {
	ID            int64               `json:"id"`             // id of country group
	CountryList   []string            `json:"country_list"`   // country group list
	PlatformGroup []PlatformGroupInfo `json:"platform_group"` // device platform list under a certain country group
}

type PlatformGroupInfo struct {
	ID       int64      `json:"id"`       // platform id
	Platform string     `json:"platform"` // device platform
	Rules    []RuleInfo `json:"rules"`    // rules under a certain platform
}

type RuleInfo struct {
	ID               int64           `json:"id"`        // rule id
	Status           bool            `json:"status"`    // enable status
	Name             string          `json:"name"`      // rule name
	Priority         int32           `json:"priority"`  // priority number
	Condition        string          `json:"condition"` // rule's condition
	DecisionInfoList []*DecisionInfo `json:"decision_info_list"`
}

type DecisionInfo struct {
	Name      *string     `json:"name,omitempty"`
	Decisions []*Decision `json:"decisions"`
}

type Decision struct {
	ID       int64  `json:"id"`       // id
	Operator string `json:"operator"` // operator
	Key      string `json:"key"`      // key
	Value    string `json:"value"`    // value
}

// 创建工单请求
type CreateTicketRequest struct {
	TenantID    int64   `json:"tenant_id"`
	Reviewer    string  `json:"reviewer"`
	Description *string `json:"description,omitempty"`
	ModuleName  string  `json:"module_name"`
	AfterRules  string  `json:"after_rules"`
	BeforeRules string  `json:"before_rules,omitempty"`
}

// 新增vod strtegy请求
type AddVodStrategyRequest struct {
	Rule *RuleGroupInfo `json:"rule"`
}

const (
	shengchao    = "shengchao"
	zhangyusf    = "yu.zhang.sf"
	shengchaoJWT = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNoZW5nY2hhb0BieXRlZGFuY2UuY29tIiwidXNlcm5hbWUiOiJzaGVuZ2NoYW8iLCJpYXQiOjE1MTYyMzkwMjJ9.mz7tCONenBTNoiOFtZCPT_MdD30No8Wfq5JAWPktoPs"
	zhangyusfJWT = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Inl1LnpoYW5nLnNmQGJ5dGVkYW5jZS5jb20iLCJ1c2VybmFtZSI6Inl1LnpoYW5nLnNmIiwiaWF0IjoxNTE2MjM5MDIyfQ.MHYHUB6XgkvNfwmxFH2eye5ki9rkQSkOgP7C5z-3qg0"
)

var name2jwt = map[string]string{
	shengchao: shengchaoJWT,
	zhangyusf: zhangyusfJWT,
}
