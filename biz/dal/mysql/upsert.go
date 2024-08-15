package mysql

import (
	"bupt/lep_user/biz/model"
	"context"

	"github.com/zhongershashen/lep_lib/dal"
	"gorm.io/gorm"
)

func UpsertPermission(ctx context.Context, db *gorm.DB, modelItem *model.LepPermissionTable) error {
	if db == nil {
		db = dal.GetDB()
	}
	query := db.Table(model.LepRoleTableName)
	return query.WithContext(ctx).Debug().Omit("created_time").Save(&modelItem).Error
}
func UpsertRolePermission(ctx context.Context, db *gorm.DB, modelItem []*model.LepRolePermissionTable) error {
	if db == nil {
		db = dal.GetDB()
	}
	query := db.Table(model.LepRolePermissionTableName)
	return query.WithContext(ctx).Debug().Omit("created_time").Save(&modelItem).Error
}
func UpsertRole(ctx context.Context, db *gorm.DB, modelItem *model.LepRoleTable) error {
	if db == nil {
		db = dal.GetDB()
	}
	query := db.Table(model.LepRoleTableName)
	return query.WithContext(ctx).Debug().Omit("created_time").Save(&modelItem).Error
}
func UpsertUser(ctx context.Context, db *gorm.DB, modelItem *model.LepUserTable) error {
	if db == nil {
		db = dal.GetDB()
	}
	query := db.Table(model.LepUserTableName)
	return query.WithContext(ctx).Debug().Omit("created_time").Save(&modelItem).Error
}
func UpsertMaterial(ctx context.Context, db *gorm.DB, modelItem []*model.LepMaterialTable) error {
	if db == nil {
		db = dal.GetDB()
	}
	query := db.Table(model.LepMaterialTableName)
	return query.WithContext(ctx).Debug().Omit("created_time").Save(&modelItem).Error
}
