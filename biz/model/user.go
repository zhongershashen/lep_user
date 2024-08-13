package model

import "time"

const LepUserTableName = "lep_user_table"

// LepUserTable 用户表
type LepUserTable struct {
	ID          int64      `json:"id" gorm:"id"`                     // 主键id
	Phone       string     `json:"phone" gorm:"phone"`               // 手机号
	Password    string     `json:"password" gorm:"password"`         // 密码，加密版
	UserName    string     `json:"user_name" gorm:"user_name"`       // 用户名
	UserAvatar  string     `json:"user_avatar" gorm:"user_avatar"`   // 用户头像
	RoleKey     string     `json:"role_key" gorm:"role_key"`         // 用户角色
	Extra       string     `json:"extra" gorm:"extra"`               // 扩展字段
	IsDeleted   int32      `json:"is_deleted" gorm:"is_deleted"`     // 0-未删除; 1-已删除
	CreatedTime *time.Time `json:"created_time" gorm:"created_time"` // 创建时间
	UpdatedTime *time.Time `json:"updated_time" gorm:"updated_time"` // 更新时间
}
