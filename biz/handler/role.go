package handler

import (
	"bupt/lep_user/biz/logic/role"
	"bupt/lep_user/kitex_gen/base"
	"bupt/lep_user/kitex_gen/lep_user"
	"context"
)

func RoleList(ctx context.Context, req *lep_user.RoleListRequest) (*lep_user.RoleListResp, error) {
	resp := new(lep_user.RoleListResp)
	resp.BaseResp = base.NewBaseResp()
	list, total, err := role.NewRoleListHandler(ctx, req).GetRoleList()
	if err != nil {
		return nil, err
	}
	resp.RoleList = list
	resp.Total = total
	return resp, err
}

//func UpsertRole(ctx context.Context, req *lep_user.UpsertRoleRequest) (*lep_user.UpsertRoleResp, error) {
//	resp := new(lep_user.UpsertPermissionResp)
//	resp.BaseResp = base.NewBaseResp()
//	handler := role.NewUpsertHandler(ctx, req)
//	err := handler.Check()
//	if err != nil {
//		return nil, err
//	}
//	id, err := handler.UpsertPermission()
//	if err != nil {
//		return nil, err
//	}
//	resp.PermissionId = id
//	return resp, err
//}
