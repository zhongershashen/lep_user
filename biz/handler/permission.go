package handler

import (
	"bupt/lep_user/kitex_gen/base"
	"bupt/lep_user/kitex_gen/lep_user"
	"context"
)

func PermissionList(ctx context.Context, req *lep_user.PermissionListRequest) (*lep_user.PermissionListResp, error) {
	resp := new(lep_user.PermissionListResp)
	resp.BaseResp = base.NewBaseResp()
	list, total, err := permission.NewActivityListHandler(ctx, req).GetPermissionList()
	return resp, err
}
