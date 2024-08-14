package mysql

import (
	"bupt/lep_user/biz/model"
	"context"

	"github.com/zhongershashen/lep_lib/dal"
	"gorm.io/gorm"
)

func CountPermission(ctx context.Context, db *gorm.DB, conditions map[string]interface{}) (total int64, err error) {
	if db == nil {
		db = dal.GetDB()
	}
	query := db.Table(model.LepPermissionTableName)
	for key, value := range conditions {
		query = query.Where(key, value)
	}
	err = query.WithContext(ctx).Debug().Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}
func QueryPermission(ctx context.Context, db *gorm.DB, conditions map[string]interface{}, offset, limit int) (modelList []*model.LepPermissionTable, err error) {
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
func CountMaterial(ctx context.Context, db *gorm.DB, conditions map[string]interface{}) (total int64, err error) {
	if db == nil {
		db = dal.GetDB()
	}
	query := db.Table(model.LepMaterialTableName)
	for key, value := range conditions {
		query = query.Where(key, value)
	}
	err = query.WithContext(ctx).Debug().Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}
func QueryMaterial(ctx context.Context, db *gorm.DB, conditions map[string]interface{}, offset, limit int) (modelList []*model.LepMaterialTable, err error) {
	if db == nil {
		db = dal.GetDB()
	}
	modelList = make([]*model.LepMaterialTable, 0)
	query := db.Table(model.LepMaterialTableName)
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
func CountRole(ctx context.Context, db *gorm.DB, conditions map[string]interface{}) (total int64, err error) {
	if db == nil {
		db = dal.GetDB()
	}
	query := db.Table(model.LepRoleTableName)
	for key, value := range conditions {
		query = query.Where(key, value)
	}
	err = query.WithContext(ctx).Debug().Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}
func QueryRole(ctx context.Context, db *gorm.DB, conditions map[string]interface{}, offset, limit int) (modelList []*model.LepRoleTable, err error) {
	if db == nil {
		db = dal.GetDB()
	}
	modelList = make([]*model.LepRoleTable, 0)
	query := db.Table(model.LepRoleTableName)
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
func QueryRolePermission(ctx context.Context, db *gorm.DB, conditions map[string]interface{}, offset, limit int) (modelList []*model.LepRolePermissionTable, err error) {
	if db == nil {
		db = dal.GetDB()
	}
	modelList = make([]*model.LepRolePermissionTable, 0)
	query := db.Table(model.LepRolePermissionTableName)
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
func CountUser(ctx context.Context, db *gorm.DB, conditions map[string]interface{}) (total int64, err error) {
	if db == nil {
		db = dal.GetDB()
	}
	query := db.Table(model.LepUserTableName)
	for key, value := range conditions {
		query = query.Where(key, value)
	}
	err = query.WithContext(ctx).Debug().Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}
func QueryUser(ctx context.Context, db *gorm.DB, conditions map[string]interface{}, offset, limit int) (modelList []*model.LepUserTable, err error) {
	if db == nil {
		db = dal.GetDB()
	}
	modelList = make([]*model.LepUserTable, 0)
	query := db.Table(model.LepUserTableName)
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
