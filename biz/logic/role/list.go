package role

import (
	"bupt/lep_user/biz/dal/mysql"
	"bupt/lep_user/kitex_gen/lep_user"
	"context"

	"github.com/zhongershashen/lep_lib/dal"
)

type ListHandler struct {
	ctx context.Context
	req *lep_user.RoleListRequest
}

func NewRoleListHandler(ctx context.Context, req *lep_user.RoleListRequest) *ListHandler {
	return &ListHandler{
		ctx: ctx,
		req: req,
	}
}
func (h *ListHandler) GetRoleList() ([]*lep_user.Role, int64, error) {
	list := make([]*lep_user.Role, 0)
	db := dal.GetDB()
	condition := h.buildCondition()
	total, err := mysql.CountRole(h.ctx, db, condition)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return list, 0, nil
	}
	modelList, err := mysql.QueryRole(h.ctx, db, condition, int(h.req.Offset), int(h.req.Limit))
	for _, item := range modelList {
		tempRole := &lep_user.Role{
			Id:             &item.ID,
			RoleName:       item.RoleName,
			RoleKey:        item.RoleKey,
			RoleDesc:       item.RoleDesc,
			Extra:          item.Extra,
			CreatedTime:    item.CreatedTime.Unix(),
			UpdatedTime:    nil,
			PermissionList: nil,
		}
		list = append(list, tempRole)
	}
}
func (h *ListHandler) buildCondition() map[string]interface{} {
	condition := make(map[string]interface{}, 0)
	req := h.req
	if req.RoleKey != nil {
		condition["role_key"] = *req.RoleKey
	}
	if req.RoleName != nil {
		condition["role_name"] = *req.RoleName
	}
	return condition
}
