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
func (h *ListHandler) GetPermissionList() ([]*lep_user.Permission, int64, error) {
	list := make([]*lep_user.Permission, 0)
	db := dal.GetDB()
	condition := h.buildCondition()
	total, err := mysql.CountPermission(h.ctx, db, condition)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return list, 0, nil
	}
	modelList, err := mysql.QueryPermission(h.ctx, db, condition, int(h.req.Offset), int(h.req.Limit))
	for _, item := range modelList {
		tempPermission := &lep_user.Permission{
			Id:             &item.ID,
			PermissionKey:  item.PermissionKey,
			PermissionName: item.PermissionName,
			PermissionDesc: item.PermissionDesc,
			Extra:          item.Extra,
			CreatedTime:    item.CreatedTime.Unix(),
			UpdatedTime:    item.UpdatedTime.Unix(),
		}
		list = append(list, tempPermission)
	}
	return list, total, nil
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
