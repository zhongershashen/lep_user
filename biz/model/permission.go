package model

import "time"

const LepPermissionTableName = "lep_permission_table"

// LepPermissionTable 系统权限
type LepPermissionTable struct {
	ID             int64      `json:"id" gorm:"id"`                           // 主键id
	PermissionName string     `json:"permission_name" gorm:"permission_name"` // 权限名称
	PermissionKey  string     `json:"permission_key" gorm:"permission_key"`   // 权限Key值
	PermissionDesc string     `json:"permission_desc" gorm:"permission_desc"` // 描述-备注
	Extra          string     `json:"extra" gorm:"extra"`                     // 扩展字段
	IsDeleted      int32      `json:"is_deleted" gorm:"is_deleted"`           // 0-未删除; 1-已删除
	CreatedTime    *time.Time `json:"created_time" gorm:"created_time"`       // 创建时间
	UpdatedTime    *time.Time `json:"updated_time" gorm:"updated_time"`       // 更新时间
}
