package permission

import (
	"bupt/lep_user/kitex_gen/lep_user"
	"context"
)

type ListHandler struct {
	ctx context.Context
	req *lep_user.PermissionListRequest
}

func NewPermissionListHandler(ctx context.Context, req *lep_user.PermissionListRequest) *ListHandler {
	return &ListHandler{
		ctx: ctx,
		req: req,
	}
}
func (h *ListHandler) GetPermissionList() (*lep_user.PermissionListResp, error) {
	resp := new(lep_user.PermissionListResp)
	list := make([]*lep_user.Permission, 0)
}
