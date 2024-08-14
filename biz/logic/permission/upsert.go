package permission

import (
	"bupt/lep_user/biz/dal/mysql"
	"bupt/lep_user/biz/model"
	"bupt/lep_user/kitex_gen/lep_user"
	"context"
	"errors"

	"github.com/zhongershashen/lep_lib/dal"
)

type UpsertHandler struct {
	ctx context.Context
	req *lep_user.UpsertPermissionRequest
}

func NewUpsertHandler(ctx context.Context, req *lep_user.UpsertPermissionRequest) *UpsertHandler {
	return &UpsertHandler{
		ctx: ctx,
		req: req,
	}
}
func (h *UpsertHandler) UpsertPermission() (int64, error) {
	db := dal.GetDB()
	if h.req.Permission == nil {
		return 0, errors.New("invalid permission")
	}
	permission := h.req.Permission
	modelItem := &model.LepPermissionTable{
		PermissionName: permission.GetPermissionName(),
		PermissionKey:  permission.PermissionKey,
		PermissionDesc: permission.GetPermissionDesc(),
	}
	if permission.Id != nil {
		modelItem.ID = *permission.Id
	}
	err := mysql.UpsertPermission(h.ctx, db, modelItem)
	if err != nil {
		return 0, err
	}
	return modelItem.ID, nil
}
func (h *UpsertHandler) Check() error {
	if h.req.Permission == nil {
		return errors.New("invalid permission")
	}
	permission := h.req.Permission
	if permission.PermissionName == "" {
		return errors.New("invalid permission_name")
	}
	if permission.PermissionKey == "" {
		return errors.New("invalid permission_key")
	}
	return nil
}
