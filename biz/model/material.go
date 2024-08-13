package model

import "time"

const LepMaterialTableName = "lep_material_table"

// LepMaterialTable 物料表
type LepMaterialTable struct {
	ID            int64      `json:"id" gorm:"id"`                         // 主键id
	UserId        int64      `json:"user_id" gorm:"user_id"`               // 用户id
	MaterialType  int8       `json:"material_type" gorm:"material_type"`   // 物料类型
	MaterialKey   string     `json:"material_key" gorm:"material_key"`     // 物料key
	MaterialValue string     `json:"material_value" gorm:"material_value"` // 物料值
	IsDeleted     int32      `json:"is_deleted" gorm:"is_deleted"`         // 0-未删除; 1-已删除
	CreatedTime   *time.Time `json:"created_time" gorm:"created_time"`     // 创建时间
	UpdatedTime   *time.Time `json:"updated_time" gorm:"updated_time"`     // 更新时间
}
