package user

import (
	"bupt/lep_user/biz/dal/mysql"
	"bupt/lep_user/kitex_gen/lep_user"
	"context"

	"github.com/zhongershashen/lep_lib/dal"
	"gorm.io/gorm"
)

type ListHandler struct {
	ctx           context.Context
	req           *lep_user.UserListRequest
	tx            *gorm.DB
	roleKeyList   []string
	permissionMap map[string][]*lep_user.Permission
}

func NewUserListHandler(ctx context.Context, req *lep_user.UserListRequest) *ListHandler {
	return &ListHandler{
		ctx: ctx,
		req: req,
	}
}
func (h *ListHandler) GetUserList() ([]*lep_user.User, int64, error) {
	list := make([]*lep_user.User, 0)
	h.roleKeyList = make([]string, 0)
	h.tx = dal.GetDB().Begin()
	condition := h.buildCondition()
	total, err := mysql.CountUser(h.ctx, h.tx, condition)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return list, 0, nil
	}
	modelList, err := mysql.QueryUser(h.ctx, h.tx, condition, int(h.req.Offset), int(h.req.Limit))
	for _, item := range modelList {
		tempUser := &lep_user.User{
			Id:          &item.ID,
			UserName:    item.UserName,
			UserAvatar:  item.UserAvatar,
			RoleKey:     item.RoleKey,
			Extra:       item.Extra,
			CreatedTime: item.CreatedTime.Unix(),
			UpdatedTime: item.UpdatedTime.Unix(),
			Phone:       item.Phone,
			RoleName:    item.RoleName,
		}
		list = append(list, tempUser)
		h.roleKeyList = append(h.roleKeyList, item.RoleKey)
	}
	err = h.loadRolePermission()
	if err != nil {
		return nil, 0, err
	}
	for _, item := range list {
		if val, ok := h.permissionMap[item.RoleKey]; ok {
			item.PermissionList = val
		}
	}
	return list, total, nil
}
func (h *ListHandler) buildCondition() map[string]interface{} {
	condition := make(map[string]interface{}, 0)
	req := h.req
	if req.RoleKey != nil {
		condition["role_key = ?"] = *req.RoleKey
	}
	if req.UserId != nil {
		condition["id = ?"] = *req.UserId
	}
	if req.Phone != nil {
		condition["phone = ?"] = *req.Phone
	}
	return condition
}
func (h *ListHandler) loadRolePermission() error {
	permissionMap := make(map[string][]*lep_user.Permission, 0)
	condition := map[string]interface{}{
		"role_key in (?)": h.roleKeyList,
	}
	permission, err := mysql.QueryRolePermission(h.ctx, h.tx, condition, 0, 0)
	if err != nil {
		return err
	}
	for _, item := range permission {
		if _, ok := permissionMap[item.RoleKey]; !ok {
			permissionMap[item.RoleKey] = make([]*lep_user.Permission, 0)
		}
		permissionMap[item.RoleKey] = append(permissionMap[item.RoleKey], &lep_user.Permission{
			PermissionKey:  item.PermissionKey,
			PermissionName: item.PermissionName,
			Extra:          item.Extra,
			CreatedTime:    item.CreatedTime.Unix(),
			UpdatedTime:    item.UpdatedTime.Unix(),
		})
	}
	h.permissionMap = permissionMap
	return nil
}
