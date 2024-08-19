package handler

import (
	"bupt/lep_user/biz/logic/user"
	"bupt/lep_user/kitex_gen/base"
	"bupt/lep_user/kitex_gen/lep_user"
	"context"
)

func UserList(ctx context.Context, req *lep_user.UserListRequest) (*lep_user.UserListResp, error) {
	resp := new(lep_user.UserListResp)
	resp.BaseResp = base.NewBaseResp()
	list, total, err := user.NewUserListHandler(ctx, req).GetUserList()
	if err != nil {
		return nil, err
	}
	resp.UserList = list
	resp.Total = total
	return resp, err
}

func UpsertUser(ctx context.Context, req *lep_user.UpsertUserRequest) (*lep_user.UpsertUserResp, error) {
	resp := new(lep_user.UpsertUserResp)
	resp.BaseResp = base.NewBaseResp()
	handler := user.NewUpsertHandler(ctx, req)
	err := handler.Check()
	if err != nil {
		return nil, err
	}
	id, err := handler.UpsertUser()
	if err != nil {
		return nil, err
	}
	resp.UserId = id
	return resp, err
}
func CheckUser(ctx context.Context, req *lep_user.CheckUserReq) (*lep_user.CheckUserResp, error) {
	resp := new(lep_user.CheckUserResp)
	resp.BaseResp = base.NewBaseResp()
	id, err := user.NewCheckHandler(ctx, req).CheckUser()
	if err != nil {
		return nil, err
	}
	resp.UserId = &id
	return resp, err
}
func DeleteUser(ctx context.Context, req *lep_user.DeleteUserRequest) (*lep_user.DeleteUserResp, error) {
	resp := new(lep_user.DeleteUserResp)
	resp.BaseResp = base.NewBaseResp()
	handler := user.NewDeleteHandler(ctx, req)
	err := handler.Delete()
	if err != nil {
		return nil, err
	}
	return resp, err
}
