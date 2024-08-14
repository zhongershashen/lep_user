package dal

import (
	"bupt/lep_backend/biz/model"
	"bupt/lep_lib/dal"
	"context"

	"gorm.io/gorm"
)

func CountPermission(ctx context.Context, db *gorm.DB, conditions map[string]interface{}) (total int64, err error) {
	if db == nil {
		db = dal.GetDB(ctx)
	}
	query := db.Table(model.LepPermissionTable)
	for key, value := range conditions {
		query = query.Where(key, value)
	}
	err = query.WithContext(ctx).Debug().Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}
func GetShBusinessActivity(ctx context.Context, db *gorm.DB, conditions map[string]interface{}, offset, limit int) (modelList []*model.LepPermissionTable, err error) {
	if db == nil {
		db = dal.GetDB(ctx)
	}
	modelList = make([]*model.LepPermissionTable, 0)
	query := db.Table(model.LepPermissionTableName)
	for key, value := range conditions {
		query = query.Where(key, value)
	}
	if offset != 0 {
		query = query.Offset(offset)
	}
	if limit != 0 {
		query = query.Limit(limit)
	}
	err = query.WithContext(ctx).Debug().Order("id desc").Find(&modelList).Error
	if err != nil {
		return nil, err
	}
	return modelList, nil
}
