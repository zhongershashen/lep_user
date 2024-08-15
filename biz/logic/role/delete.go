package role

import (
	"bupt/lep_user/biz/dal/mysql"
	"bupt/lep_user/kitex_gen/lep_user"
	"context"
	"errors"

	"github.com/zhongershashen/lep_lib/dal"
	"gorm.io/gorm"
)

type DeleteHandler struct {
	ctx context.Context
	req *lep_user.DeleteRoleRequest
	tx  *gorm.DB
}

func NewDeleteHandler(ctx context.Context, req *lep_user.DeleteRoleRequest) *DeleteHandler {
	return &DeleteHandler{
		ctx: ctx,
		req: req,
	}
}
func (h *DeleteHandler) Delete() error {
	var err error
	h.tx = dal.GetDB().Begin()
	defer func() {
		if err != nil {
			h.tx.Rollback()
		} else {
			h.tx.Commit()
		}
	}()
	err = h.deleteMainTable()
	if err != nil {
		return err
	}
	err = h.deleteRolePermission()
	if err != nil {
		return err
	}
	return nil
}
func (h *DeleteHandler) deleteMainTable() error {
	req := h.req
	if req.Id == 0 {
		return errors.New("invalid id")
	}
	condition := map[string]interface{}{
		"id = ?": req.Id,
	}
	updates := map[string]interface{}{
		"is_deleted": 1,
	}
	err := mysql.UpdateRole(h.ctx, h.tx, condition, updates)
	if err != nil {
		return err
	}
	return nil
}
func (h *DeleteHandler) deleteRolePermission() error {
	req := h.req
	if req.Id == 0 {
		return errors.New("invalid id")
	}
	condition := map[string]interface{}{
		"id = ?": req.Id,
	}
	role, err := mysql.QueryRole(h.ctx, h.tx, condition, 0, 0)
	if err != nil {
		return err
	}

	roleKey := role[0].RoleKey
	condition = map[string]interface{}{
		"role_key = ?": roleKey,
	}
	updates := map[string]interface{}{
		"is_deleted": 1,
	}
	err = mysql.UpdateRolePermission(h.ctx, h.tx, condition, updates)
	if err != nil {
		return err
	}
	return nil
}
