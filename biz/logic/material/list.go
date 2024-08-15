package material

import (
	"bupt/lep_user/biz/dal/mysql"
	"bupt/lep_user/kitex_gen/lep_user"
	"context"

	"github.com/zhongershashen/lep_lib/dal"
)

type ListHandler struct {
	ctx context.Context
	req *lep_user.MaterialListRequest
}

func NewMaterialListHandler(ctx context.Context, req *lep_user.MaterialListRequest) *ListHandler {
	return &ListHandler{
		ctx: ctx,
		req: req,
	}
}
func (h *ListHandler) GetMaterialList() ([]*lep_user.Material, int64, error) {
	list := make([]*lep_user.Material, 0)
	db := dal.GetDB()
	condition := h.buildCondition()
	total, err := mysql.CountMaterial(h.ctx, db, condition)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return list, 0, nil
	}
	modelList, err := mysql.QueryMaterial(h.ctx, db, condition, int(h.req.Offset), int(h.req.Limit))
	for _, item := range modelList {
		tempMaterial := &lep_user.Material{
			Id:            item.ID,
			MaterialKey:   item.MaterialKey,
			MaterialType:  int64(item.MaterialType),
			CreatedTime:   item.CreatedTime.Unix(),
			UpdatedTime:   item.UpdatedTime.Unix(),
			MaterialValue: item.MaterialValue,
		}
		list = append(list, tempMaterial)
	}
	return list, total, nil
}
func (h *ListHandler) buildCondition() map[string]interface{} {
	condition := make(map[string]interface{}, 0)
	req := h.req
	if req.MaterialKey != nil {
		condition["material_key = ?"] = *req.MaterialKey
	}
	if req.MaterialType != nil {
		condition["material_type = ?"] = *req.MaterialType
	}
	if req.UserId != nil {
		condition["user_id = ?"] = *req.UserId
	}
	return condition
}
