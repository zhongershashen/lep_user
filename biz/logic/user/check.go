package user

import (
	"bupt/lep_user/biz/dal/mysql"
	"bupt/lep_user/kitex_gen/lep_user"
	"context"

	"github.com/zhongershashen/lep_lib/dal"
)

type CheckHandler struct {
	ctx context.Context
	req *lep_user.CheckUserReq
}

func NewCheckHandler(ctx context.Context, req *lep_user.CheckUserReq) *CheckHandler {
	return &CheckHandler{
		ctx: ctx,
		req: req,
	}
}
func (h *CheckHandler) CheckUser() (int64, error) {
	var err error
	db := dal.GetDB()
	if err != nil {
		return 0, err
	}
	req := h.req
	condition := make(map[string]interface{}, 0)
	if req.Password != "" && req.UserName != "" {
		condition["pass_word = ?"] = req.Password
		condition["user_name = ?"] = req.UserName
	}
	if req.Phone != "" {
		condition["phone = ?"] = req.Phone
	}
	modelList, err := mysql.QueryUser(h.ctx, db, condition, 0, 1)
	if err != nil {
		return 0, err
	}
	if len(modelList) > 0 {
		return modelList[0].ID, nil
	}
	return 0, nil
}
