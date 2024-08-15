package mysql

import (
	"bupt/lep_user/biz/model"
	"context"

	"github.com/zhongershashen/lep_lib/dal"
	"gorm.io/gorm"
)

func UpdatePermission(ctx context.Context, db *gorm.DB, conditions, updates map[string]interface{}) error {
	if db == nil {
		db = dal.GetDB()
	}
	query := db.Table(model.LepMaterialTableName)
	for key, value := range conditions {
		query = query.Where(key, value)
	}
	return query.WithContext(ctx).Debug().Updates(updates).Error
}
func UpdateRole(ctx context.Context, db *gorm.DB, conditions, updates map[string]interface{}) error {
	if db == nil {
		db = dal.GetDB()
	}
	query := db.Table(model.LepRoleTableName)
	for key, value := range conditions {
		query = query.Where(key, value)
	}
	return query.WithContext(ctx).Debug().Updates(updates).Error
}
func UpdateRolePermission(ctx context.Context, db *gorm.DB, conditions, updates map[string]interface{}) error {
	if db == nil {
		db = dal.GetDB()
	}
	query := db.Table(model.LepRolePermissionTableName)
	for key, value := range conditions {
		query = query.Where(key, value)
	}
	return query.WithContext(ctx).Debug().Updates(updates).Error
}
func UpdateUser(ctx context.Context, db *gorm.DB, conditions, updates map[string]interface{}) error {
	if db == nil {
		db = dal.GetDB()
	}
	query := db.Table(model.LepUserTableName)
	for key, value := range conditions {
		query = query.Where(key, value)
	}
	return query.WithContext(ctx).Debug().Updates(updates).Error
}
func UpdateMaterial(ctx context.Context, db *gorm.DB, conditions, updates map[string]interface{}) error {
	if db == nil {
		db = dal.GetDB()
	}
	query := db.Table(model.LepMaterialTableName)
	for key, value := range conditions {
		query = query.Where(key, value)
	}
	return query.WithContext(ctx).Debug().Updates(updates).Error
}
