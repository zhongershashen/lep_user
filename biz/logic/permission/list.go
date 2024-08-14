package permission

import (
	"bupt/lep_user/biz/dal/mysql"
	"bupt/lep_user/kitex_gen/lep_user"
	"context"

	"github.com/zhongershashen/lep_lib/dal"
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
	db := dal.GetDB()
	condition := h.buildCondition()
	total, err := mysql.CountPermission(h.ctx, db, condition)
	if err != nil {
		return nil, err
	}
	if total == 0 {
		resp.Total = 0
		resp.PermissionList = list
		return resp, nil
	}
	list, err = mysql.QueryPermission(h.ctx, db, condition)
}
func (h *ListHandler) buildCondition() map[string]interface{} {
	condition := make(map[string]interface{}, 0)
	req := h.req
	if req.PermissionKey != nil {
		condition["permission_key"] = *req.PermissionKey
	}
	if req.PermissionName != nil {
		condition["permission_name"] = *req.PermissionName
	}
	return condition
}
