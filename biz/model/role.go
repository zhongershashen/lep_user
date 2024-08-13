package model

import "time"

const (
	LepRoleTableName           = "lep_role_table"
	LepRolePermissionTableName = "lep_role_permission_table"
)

// LepRoleTable 角色表
type LepRoleTable struct {
	ID          int64      `json:"id" gorm:"id"`                     // 主键id
	RoleName    string     `json:"role_name" gorm:"role_name"`       // 角色名称
	RoleKey     string     `json:"role_key" gorm:"role_key"`         // 角色key
	RoleDesc    string     `json:"role_desc" gorm:"role_desc"`       // 描述-备注
	Extra       string     `json:"extra" gorm:"extra"`               // 扩展字段
	IsDeleted   int32      `json:"is_deleted" gorm:"is_deleted"`     // 0-未删除; 1-已删除
	CreatedTime *time.Time `json:"created_time" gorm:"created_time"` // 创建时间
	UpdatedTime *time.Time `json:"updated_time" gorm:"updated_time"` // 更新时间
}

// LepRolePermissionTable 角色-权限表
type LepRolePermissionTable struct {
	ID            int64      `json:"id" gorm:"id"`                         // 主键id
	RoleKey       string     `json:"role_key" gorm:"role_key"`             // 角色key
	PermissionKey string     `json:"permission_key" gorm:"permission_key"` // 权限key
	Extra         string     `json:"extra" gorm:"extra"`                   // 扩展字段
	IsDeleted     int32      `json:"is_deleted" gorm:"is_deleted"`         // 0-未删除; 1-已删除
	CreatedTime   *time.Time `json:"created_time" gorm:"created_time"`     // 创建时间
	UpdatedTime   *time.Time `json:"updated_time" gorm:"updated_time"`     // 更新时间
}
