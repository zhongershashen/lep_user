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
	resp, err = handler.UpsertRole(ctx, req)
	return
}

// UserList implements the LepUserImpl interface.
func (s *LepUserImpl) UserList(ctx context.Context, req *lep_user.UserListRequest) (resp *lep_user.UserListResp, err error) {
	resp, err = handler.UserList(ctx, req)
	return
}

// UpsertUser implements the LepUserImpl interface.
func (s *LepUserImpl) UpsertUser(ctx context.Context, req *lep_user.UpsertUserRequest) (resp *lep_user.UpsertUserResp, err error) {
	resp, err = handler.UpsertUser(ctx, req)
	return
}

// MaterialList implements the LepUserImpl interface.
func (s *LepUserImpl) MaterialList(ctx context.Context, req *lep_user.MaterialListRequest) (resp *lep_user.MaterialListResp, err error) {
	resp, err = handler.MaterialList(ctx, req)
	return
}

// UpsertMaterial implements the LepUserImpl interface.
func (s *LepUserImpl) UpsertMaterial(ctx context.Context, req *lep_user.UpsertMaterialRequest) (resp *lep_user.UpsertMaterialResp, err error) {
	resp, err = handler.UpsertMaterial(ctx, req)
	return
}

// DeletePermission implements the LepUserImpl interface.
func (s *LepUserImpl) DeletePermission(ctx context.Context, req *lep_user.DeletePermissionRequest) (resp *lep_user.DeletePermissionResp, err error) {
	resp, err = handler.DeletePermission(ctx, req)
	return
}

// DeleteRole implements the LepUserImpl interface.
func (s *LepUserImpl) DeleteRole(ctx context.Context, req *lep_user.DeleteRoleRequest) (resp *lep_user.DeleteRoleResp, err error) {
	resp, err = handler.DeleteRole(ctx, req)
	return
}

// DeleteUser implements the LepUserImpl interface.
func (s *LepUserImpl) DeleteUser(ctx context.Context, req *lep_user.DeleteUserRequest) (resp *lep_user.DeleteUserResp, err error) {
	resp, err = handler.DeleteUser(ctx, req)
	return
}

// DeleteMaterial implements the LepUserImpl interface.
func (s *LepUserImpl) DeleteMaterial(ctx context.Context, req *lep_user.DeleteMaterialRequest) (resp *lep_user.DeleteMaterialResp, err error) {
	resp, err = handler.DeleteMaterial(ctx, req)
	return
}
