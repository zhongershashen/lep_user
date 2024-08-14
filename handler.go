package main

import (
	"bupt/lep_user/biz/handler"
	"bupt/lep_user/kitex_gen/lep_user"
	"context"
)

// LepUserImpl implements the last service interface defined in the IDL.
type LepUserImpl struct{}

// PermissionList implements the LepUserImpl interface.
func (s *LepUserImpl) PermissionList(ctx context.Context, req *lep_user.PermissionListRequest) (resp *lep_user.PermissionListResp, err error) {
	resp, err = handler.PermissionList(ctx, req)
	return
}

// UpsertPermission implements the LepUserImpl interface.
func (s *LepUserImpl) UpsertPermission(ctx context.Context, req *lep_user.UpsertPermissionRequest) (resp *lep_user.UpsertPermissionResp, err error) {
	resp, err = handler.UpsertPermission(ctx, req)
	return
}

// RoleList implements the LepUserImpl interface.
func (s *LepUserImpl) RoleList(ctx context.Context, req *lep_user.RoleListRequest) (resp *lep_user.RoleListResp, err error) {
	resp, err = handler.RoleList(ctx, req)
	return
}

// UpsertRole implements the LepUserImpl interface.
func (s *LepUserImpl) UpsertRole(ctx context.Context, req *lep_user.UpsertRoleRequest) (resp *lep_user.UpsertRoleResp, err error) {
	// TODO: Your code here...
	return
}

// UserList implements the LepUserImpl interface.
func (s *LepUserImpl) UserList(ctx context.Context, req *lep_user.UserListRequest) (resp *lep_user.UserListResp, err error) {
	// TODO: Your code here...
	return
}

// UpsertUser implements the LepUserImpl interface.
func (s *LepUserImpl) UpsertUser(ctx context.Context, req *lep_user.UpsertUserRequest) (resp *lep_user.UpsertUserResp, err error) {
	// TODO: Your code here...
	return
}

// MaterialList implements the LepUserImpl interface.
func (s *LepUserImpl) MaterialList(ctx context.Context, req *lep_user.MaterialListRequest) (resp *lep_user.MaterialListResp, err error) {
	// TODO: Your code here...
	return
}

// UpsertMaterial implements the LepUserImpl interface.
func (s *LepUserImpl) UpsertMaterial(ctx context.Context, req *lep_user.UpsertMaterialRequest) (resp *lep_user.UpsertMaterialResp, err error) {
	// TODO: Your code here...
	return
}
