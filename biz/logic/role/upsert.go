package role

import (
	"bupt/lep_user/biz/dal/mysql"
	"bupt/lep_user/biz/model"
	"bupt/lep_user/kitex_gen/lep_user"
	"context"
	"errors"

	"github.com/zhongershashen/lep_lib/utils"

	"github.com/zhongershashen/lep_lib/dal"
	"gorm.io/gorm"
)

type UpsertHandler struct {
	ctx context.Context
	req *lep_user.UpsertRoleRequest
	tx  *gorm.DB
	id  int64
}

func NewUpsertHandler(ctx context.Context, req *lep_user.UpsertRoleRequest) *UpsertHandler {
	return &UpsertHandler{
		ctx: ctx,
		req: req,
	}
}
func (h *UpsertHandler) UpsertRole() (int64, error) {
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
	err = h.upsertPermission()
	return h.id, nil
}
func (h *UpsertHandler) Check() error {
	if h.req.Role == nil {
		return errors.New("invalid role")
	}
	role := h.req.Role
	if role.RoleName == "" {
		return errors.New("invalid role_name")
	}
	if role.RoleKey == "" {
		return errors.New("invalid role_key")
	}
	return nil
}
func (h *UpsertHandler) upsertMainTable() error {
	role := h.req.Role
	modelItem := &model.LepRoleTable{
		RoleName:  role.GetRoleName(),
		RoleKey:   role.GetRoleKey(),
		RoleDesc:  role.GetRoleDesc(),
		Extra:     role.Extra,
		IsDeleted: 0,
	}
	if role.Id != nil {
		modelItem.ID = *role.Id
	}
	err := mysql.UpsertRole(h.ctx, h.tx, modelItem)
	if err != nil {
		return err
	}
	h.id = modelItem.ID
	return nil
}

func (h *UpsertHandler) upsertPermission() error {
	role := h.req.Role
	var (
		permissionList []string
	)
	conditions := map[string]interface{}{
		"role_key = ?":   h.id,
		"is_deleted = ?": 0,
	}
	modelList, err := mysql.QueryRolePermission(h.ctx, h.tx, conditions, 0, 0)
	if err != nil {
		return err
	}
	for _, item := range modelList {
		permissionList = append(permissionList, item.PermissionKey)
	}
	for _, item := range role.PermissionList { // 新增sku
		if item == nil {
			continue
		}
		if !utils.ContainsString(permissionList, item.PermissionKey) {
			modelList = append(modelList, &model.LepRolePermissionTable{
				RoleKey:        role.RoleKey,
				RoleName:       role.RoleName,
				PermissionName: item.PermissionName,
				PermissionKey:  item.PermissionKey,
				Extra:          item.Extra,
				IsDeleted:      0,
			})
		}
	}
	for _, item := range modelList { // 删除
		if !utils.ContainsString(permissionList, item.PermissionKey) {
			item.IsDeleted = 1
		}
	}
	err = mysql.UpsertRolePermission(h.ctx, h.tx, modelList)
	if err != nil {
		return err
	}
	return nil
}
