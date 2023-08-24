package gorm

import (
	"time"

	"gorm.io/gorm"
)

type DBAppInfo struct {
	ID        int64     `gorm:"column:id" json:"-"`                // 主键，自增长
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`        // 创建时间，自动填充
	UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`        // 更新时间，自动填充
	AppID     int       `gorm:"column:app_id" json:"app_id"`       // the id of app [for request/response]
	AppName   string    `gorm:"column:app_name" json:"app_name"`   // the name of app [for request/response]
	Creator   string    `gorm:"column:creator" json:"creator"`     // the name of creator who created it [for request/response]
	IconPath  string    `gorm:"column:icon_path" json:"icon_path"` // the path of icon file(this path is a accessible path) [for request/response]
	Admins    string    `gorm:"column:admins" json:"admins"`       // the administers of this app including creator by default [for request/response]
}

func (d *DBAppInfo) TableName() string {
	return "app_info"
}

type DBModuleInfo struct {
	ID         int64     `gorm:"column:id" json:"-"`                    // 主键，自增长
	CreatedAt  time.Time `gorm:"column:created_at" json:"-"`            // 创建时间，自动填充
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"-"`            // 更新时间，自动填充
	AppID      int       `gorm:"column:app_id" json:"app_id"`           // app id
	Creator    string    `gorm:"column:creator" json:"creator"`         // always "built-in"
	IconPath   string    `gorm:"column:icon_path" json:"icon_path"`     // using default icon file if empty
	ModuleName string    `gorm:"column:module_name" json:"module_name"` // module name
}

func (d *DBModuleInfo) TableName() string {
	return "module_info"
}

type DBTenantInfo struct {
	gorm.Model
	ID         int64     `gorm:"column:id" json:"-"`                    // 主键，自增长
	CreatedAt  time.Time `gorm:"column:created_at" json:"-"`            // 创建时间，自动填充
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"-"`            // 更新时间，自动填充
	AppID      int       `gorm:"column:app_id" json:"app_id"`           // the id of app [for request]
	TenantID   int       `gorm:"column:tenant_id" json:"tenant_id"`     // the id of tenant [for request/response]
	ModuleName string    `gorm:"column:module_name" json:"module_name"` // the id of module [for request]
	TenantName string    `gorm:"column:tenant_name" json:"tenant_name"` // the name of tenant [for request/response]
	Creator    string    `gorm:"column:creator" json:"creator"`         // the name of creator who created it [for request/response]
	IconPath   string    `gorm:"column:icon_path" json:"icon_path"`     // the path of accessible icon file [for request/response]
	Admins     string    `gorm:"column:admins" json:"admins"`           // the admins of this tenant included creator by default [for request/response]
	ReadOnlys  string    `gorm:"column:read_onlys" json:"read_onlys"`   // read-only users of this tenant [for request/response]
}

func (d *DBTenantInfo) TableName() string {
	return "tenant_info"
}
