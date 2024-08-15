package user

import (
	"bupt/lep_user/biz/dal/mysql"
	"bupt/lep_user/kitex_gen/lep_user"
	"context"
	"errors"

	"github.com/zhongershashen/lep_lib/dal"
)

type DeleteHandler struct {
	ctx context.Context
	req *lep_user.DeleteUserRequest
}

func NewDeleteHandler(ctx context.Context, req *lep_user.DeleteUserRequest) *DeleteHandler {
	return &DeleteHandler{
		ctx: ctx,
		req: req,
	}
}
func (h *DeleteHandler) Delete() error {
	db := dal.GetDB()
	req := h.req
	if req.Id == 0 {
		return errors.New("invalid id")
	}
	conditions := map[string]interface{}{
		"id = ?": req.Id,
	}
	updates := map[string]interface{}{
		"is_deleted": 1,
	}
	err := mysql.UpdateUser(h.ctx, db, conditions, updates)
	if err != nil {
		return err
	}
	return nil
}
