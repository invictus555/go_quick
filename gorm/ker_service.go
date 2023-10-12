package gorm

import (
	"time"
)

type DBAppInfo struct {
	ID        uint64    `gorm:"column:id;primaryKey;comment:'主键'" json:"-"` // 主键
	CreatedAt time.Time `gorm:"column:created_at;comment:'创建时间'" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at;comment:'更新时间'" json:"-"`
	AppID     int32     `gorm:"column:app_id;not null;comment:'应用ID'" json:"app_id"`
	Admins    string    `gorm:"column:admins;not null;type:varchar(2048);comment:'管理员列表'" json:"admins"`
	AppName   string    `gorm:"column:app_name;not null;type:varchar(64);comment:'应用名称'" json:"app_name"`
	Creator   string    `gorm:"column:creator;not null;type:varchar(64);comment:'创建者名称'" json:"creator"`
	IconPath  string    `gorm:"column:icon_path;type:varchar(1024);comment:'icon路径'" json:"icon_path"`
}

func (d *DBAppInfo) TableName() string {
	return "global_t_app_info"
}

type DBModuleInfo struct {
	ID         uint64    `gorm:"column:id;primaryKey;comment:'主键'" json:"-"` // 主键
	CreatedAt  time.Time `gorm:"column:created_at;comment:'创建时间'" json:"-"`
	UpdatedAt  time.Time `gorm:"column:updated_at;comment:'更新时间'" json:"-"`
	AppID      int32     `gorm:"column:app_id;not null;comment:'应用名称'" json:"app_id"`
	Creator    string    `gorm:"column:creator;not null;type:varchar(64);comment:'创建者'" json:"creator"`
	IconPath   string    `gorm:"column:icon_path;type:varchar(1024);comment:'icon路径'" json:"icon_path"`
	ModuleName string    `gorm:"column:module_name;not null;type:varchar(32);comment:'模块名称'" json:"module_name"`
}

func (d *DBModuleInfo) TableName() string {
	return "global_t_module_info"
}

// ID充当TenantID, 具有唯一性，后续操作直接可用ID充当AppID+ModuleName+TenantID的效果
type DBTenantInfo struct {
	ID         uint64    `gorm:"column:id;primaryKey;comment:'主键'" json:"-"` // 主键
	CreatedAt  time.Time `gorm:"column:created_at;comment:'创建时间'" json:"-"`
	UpdatedAt  time.Time `gorm:"column:updated_at;comment:'更新时间'" json:"-"`
	AppID      int32     `gorm:"column:app_id;not null;comment:'应用名称'" json:"app_id"`
	Admins     string    `gorm:"column:admins;not null;type:varchar(2048);comment:'管理员列表'" json:"admins"`
	Creator    string    `gorm:"column:creator;not null;type:varchar(64);comment:'创建者名称'" json:"creator"`
	TenantKey  string    `gorm:"column:tenant_key;not null;type:varchar(32);comment:'租户key'" json:"tenant_key"`
	IconPath   string    `gorm:"column:icon_path;type:varchar(1024);comment:'icon路径'" json:"icon_path"`
	ReadOnlys  string    `gorm:"column:read_onlys;type:varchar(2048);comment:'只读用户列表'" json:"read_onlys"`
	ModuleName string    `gorm:"column:module_name;not null;type:varchar(32);comment:'模块名称'" json:"module_name"`
	TenantName string    `gorm:"column:tenant_name;not null;type:varchar(64);comment:'租户名称'" json:"tenant_name"`
	ExtraInfos string    `gorm:"column:extra_infos;type:varchar(2048);comment:'额外信息'" json:"extra_infos"`
}

func (d *DBTenantInfo) TableName() string {
	return "global_t_tenant_info"
}

// 直接只用tenantID划分租户，等效于appID+moduleName+tenantID(之前创建租户信息时系统生成的ID具有唯一性)
// ParamCategory 在数据中表现为uint8
type DBParameterInfo struct {
	ID          uint64    `gorm:"column:id;primaryKey;comment:'主键'" josn:"-"` // 主键
	CreatedAt   time.Time `gorm:"column:created_at;comment:'创建时间'" json:"-"`
	UpdatedAt   time.Time `gorm:"column:updated_at;comment:'更新时间'" json:"-"`
	TenantID    uint64    `gorm:"column:tenant_id;not null;comment:'租户ID'" json:"tenant_id"` // DBTenantInfo.ID
	Creator     string    `gorm:"column:creator;not null;type:varchar(32);comment:'创建者名称'" json:"creator"`
	ParamName   string    `gorm:"column:param_name;not null;type:varchar(64);comment:'参数名'" json:"param_name"`
	ParamTag    string    `gorm:"column:param_tag;not null;type:varchar(32);comment:'参数标签'" json:"param_tag"`
	ParamType   string    `gorm:"column:param_type;not null;type:varchar(8);comment:'参数类型'" json:"param_type"`
	Category    string    `gorm:"column:category;not null;type:varchar(16);comment:'参数分类'" json:"category"` // bussiness:业务参数, common:通用参数
	Description string    `gorm:"column:description;type:varchar(256);comment:'备注描述'" json:"description"`
}

func (d *DBParameterInfo) TableName() string {
	return "global_t_parameter_info"
}

// tenantID 用于划分租户，等效于appID+moduleID+tenantID(之前创建租户信息时系统生成的ID具有唯一性)
type DBVariableInfo struct {
	ID             uint64    `gorm:"column:id;primaryKey;comment:'主键'" json:"id"` // 主键
	CreatedAt      time.Time `gorm:"column:created_at;comment:'创建时间'" json:"-"`
	UpdatedAt      time.Time `gorm:"column:updated_at;comment:'更新时间'" json:"-"`
	TenantID       uint64    `gorm:"column:tenant_id;not null;comment:'租户ID'" json:"tenant_id"` // DBTenantInfo.ID
	Creator        string    `gorm:"column:creator;not null;type:varchar(64);comment:'创建者名称'" json:"creator"`
	VarName        string    `gorm:"column:var_name;not null;type:varchar(32);comment:'变量名称'" json:"var_name"`
	VarType        string    `gorm:"column:var_type;not null;type:varchar(8);comment:'变量类型,4选1'" json:"var_type"`
	Description    string    `gorm:"column:description;type:varchar(256);comment:'对变量的描述'" json:"description"`
	Classification string    `gorm:"column:classification;not null;type:varchar(16);comment:'变量种类,common/business'" json:"classification"`
}

func (d *DBVariableInfo) TableName() string {
	return "global_t_variable_info"
}

type DBVodRuleGroup struct {
	ID          uint64    `gorm:"column:id;primaryKey;comment:'主键,自增长'"` // 主键
	GroupID     string    `gorm:"column:group_id;type:varchar(36);comment:'规则组ID'"`
	CreatedAt   time.Time `gorm:"column:created_at;comment:'创建时间,自动填充'"`
	UpdatedAt   time.Time `gorm:"column:updated_at;comment:'更新时间,自动填充'"`
	TenantID    uint64    `gorm:"column:tenant_id;comment:'租户ID,以租户划分规则归属'"`
	Creator     string    `gorm:"column:creator;type:varchar(64);comment:'规则创建者,邮箱前缀'"`
	PageName    string    `gorm:"column:page_name;type:varchar(128);comment:'这页规则的组名称'"`
	Description string    `gorm:"column:description;type:varchar(256);comment:'规则组的备注说明'"`
}

func (d *DBVodRuleGroup) TableName() string {
	return "vod_t_rule_group"
}

type DBCountryGroup struct {
	ID        uint64    `gorm:"column:id;primaryKey;comment:'主键,自增长'"` // 主键
	CreatedAt time.Time `gorm:"column:created_at;comment:'创建时间,自动填充'"`
	UpdatedAt time.Time `gorm:"column:updated_at;comment:'更新时间,自动填充'"`
	GroupID   string    `gorm:"column:group_id;type:varchar(36);comment:'关联DBVodRulePage.GroupID'"`
	Countries string    `gorm:"column:countries;type:varchar(2048);comment:'国家组,各国家分号分隔组成一行'"`
}

func (d *DBCountryGroup) TableName() string {
	return "vod_t_country_group"
}

type DBPlatform struct {
	ID             uint64    `gorm:"column:id;primaryKey;comment:'主键,自增长'"` // 主键
	CreatedAt      time.Time `gorm:"column:created_at;comment:'创建时间,自动填充'"`
	UpdatedAt      time.Time `gorm:"column:updated_at;comment:'更新时间,自动填充'"`
	Platform       string    `gorm:"column:platform;type:varchar(16);comment:'设备类型,Android IOS All 三选一'"`
	GroupID        string    `gorm:"column:group_id;type:varchar(36);comment:'关联DBVodRulePage.GroupID'"`
	CountryGroupID uint64    `gorm:"column:coontry_group_id;comment:'用于关联DBCountryGroup.ID'"`
}

func (d *DBPlatform) TableName() string {
	return "vod_t_platform"
}

type DBVodRule struct {
	ID             uint64    `gorm:"column:id;primaryKey;commemt:'主键,自增长'"` // 主键
	CreatedAt      time.Time `gorm:"column:created_at;comment:'创建时间,自动填充'"`
	UpdatedAt      time.Time `gorm:"column:updated_at;comment:'更新时间,自动填充'"`
	Status         bool      `gorm:"column:status;comment:'是否处于启用状态'"`
	Priority       uint8     `gorm:"column:priority;comment:'规则优先级,255个等级满足需求'"`
	Expression     string    `gorm:"column:express;type:varchar(2048);comment:'规则左边部分'"`
	Decision       string    `gorm:"column:decision;type:varchar(2048);comment:'规则右边部分'"`
	GroupID        string    `gorm:"column:group_id;type:varchar(36);comment:'关联DBVodRulePage.GroupID'"`
	PlatformID     uint64    `gorm:"column:platform_id;comment:'用于关联DBPlatform.ID'"`
	CountryGroupID uint64    `gorm:"column:coontry_group_id;comment:'用于关联DBCountryGroup.ID'"`
}

func (d *DBVodRule) TableName() string {
	return "vod_t_rule"
}

func TestifyCreateTable() {
	// framework
	GetDB().Set("gorm:table_options", "COMMENT='应用信息表'").AutoMigrate(&DBAppInfo{})
	GetDB().Set("gorm:table_options", "COMMENT='模块信息表'").AutoMigrate(&DBModuleInfo{})
	GetDB().Set("gorm:table_options", "COMMENT='租户信息表'").AutoMigrate(&DBTenantInfo{})
	// parameter
	GetDB().Set("gorm:table_options", "COMMENT='参数信息表'").AutoMigrate(&DBParameterInfo{})
	// variable
	GetDB().Set("gorm:table_options", "COMMENT='变量信息表'").AutoMigrate(&DBVariableInfo{})

	GetDB().Set("gorm:table_options", "COMMENT='规则组信息表'").AutoMigrate(&DBVodRuleGroup{})
	GetDB().Set("gorm:table_options", "COMMENT='国家组信息表'").AutoMigrate(&DBCountryGroup{})
	GetDB().Set("gorm:table_options", "COMMENT='platfomr信息表'").AutoMigrate(&DBPlatform{})
	GetDB().Set("gorm:table_options", "COMMENT='规则信息表'").AutoMigrate(&DBVodRule{})

}
