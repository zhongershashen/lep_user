package user

import (
	"bupt/lep_user/biz/dal/mysql"
	"bupt/lep_user/biz/model"
	"bupt/lep_user/kitex_gen/lep_user"
	"context"
	"errors"

	"github.com/zhongershashen/lep_lib/dal"
	"gorm.io/gorm"
)

type UpsertHandler struct {
	ctx context.Context
	req *lep_user.UpsertUserRequest
	tx  *gorm.DB
	id  int64
}

func NewUpsertHandler(ctx context.Context, req *lep_user.UpsertUserRequest) *UpsertHandler {
	return &UpsertHandler{
		ctx: ctx,
		req: req,
	}
}
func (h *UpsertHandler) UpsertUser() (int64, error) {
	var err error
	h.tx = dal.GetDB().Begin()
	defer func() {
		if err != nil {
			h.tx.Rollback()
		} else {
			h.tx.Commit()
		}
	}()
	err = h.upsertMainTable()
	if err != nil {
		return 0, err
	}
	return h.id, nil
}
func (h *UpsertHandler) Check() error {
	if h.req.User == nil {
		return errors.New("invalid role")
	}
	user := h.req.User
	if user.UserName == "" {
		return errors.New("invalid role_name")
	}
	if user.UserName == "" {
		return errors.New("invalid user_name")
	}
	if user.UserAvatar == "" {
		return errors.New("invalid user_avatar")
	}
	return nil
}
func (h *UpsertHandler) upsertMainTable() error {
	user := h.req.User
	// todo encrypt
	modelItem := &model.LepUserTable{
		Phone:      user.Phone,
		Password:   user.Password,
		UserName:   user.UserName,
		UserAvatar: user.UserAvatar,
		RoleKey:    user.RoleKey,
		Extra:      user.Extra,
		RoleName:   user.RoleName,
	}
	if user.Id != nil {
		modelItem.ID = *user.Id
	}
	err := mysql.UpsertUser(h.ctx, h.tx, modelItem)
	if err != nil {
		return err
	}
	h.id = modelItem.ID
	return nil
}
