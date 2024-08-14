package handler

import (
	"bupt/lep_user/biz/logic/permission"
	"bupt/lep_user/kitex_gen/base"
	"bupt/lep_user/kitex_gen/lep_user"
	"context"
)

func PermissionList(ctx context.Context, req *lep_user.PermissionListRequest) (*lep_user.PermissionListResp, error) {
	resp := new(lep_user.PermissionListResp)
	resp.BaseResp = base.NewBaseResp()
	list, total, err := permission.NewPermissionListHandler(ctx, req).GetPermissionList()
	if err != nil {
		return nil, err
	}
	resp.PermissionList = list
	resp.Total = total
	return resp, err
}
func UpsertPermission(ctx context.Context, req *lep_user.UpsertPermissionRequest) (*lep_user.UpsertPermissionResp, error) {
	resp := new(lep_user.UpsertPermissionResp)
	resp.BaseResp = base.NewBaseResp()
	handler := permission.NewUpsertHandler(ctx, req)
	err := handler.Check()
	if err != nil {
		return nil, err
	}
	id, err := handler.UpsertPermission()
	if err != nil {
		return nil, err
	}
	resp.PermissionId = id
	return resp, err
}
