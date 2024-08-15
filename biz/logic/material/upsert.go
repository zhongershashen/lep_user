package material

import (
	"bupt/lep_user/biz/dal/mysql"
	"bupt/lep_user/biz/model"
	"bupt/lep_user/kitex_gen/lep_user"
	"context"
	"errors"

	"github.com/zhongershashen/lep_lib/dal"
)

type UpsertHandler struct {
	ctx context.Context
	req *lep_user.UpsertMaterialRequest
}

func NewUpsertHandler(ctx context.Context, req *lep_user.UpsertMaterialRequest) *UpsertHandler {
	return &UpsertHandler{
		ctx: ctx,
		req: req,
	}
}
func (h *UpsertHandler) UpsertMaterial() ([]int64, error) {
	var idList []int64
	db := dal.GetDB()
	material := h.req.Material
	modelItems := make([]*model.LepMaterialTable, 0)
	for _, item := range material {
		modelItem := &model.LepMaterialTable{
			MaterialType:  int8(item.MaterialType),
			MaterialKey:   item.MaterialKey,
			MaterialValue: item.MaterialValue,
			UserId:        item.UserId,
		}
		if item.Id != 0 {
			modelItem.ID = item.Id
		}
		modelItems = append(modelItems, modelItem)
	}
	err := mysql.UpsertMaterial(h.ctx, db, modelItems)
	if err != nil {
		return nil, err
	}
	for _, item := range modelItems {
		idList = append(idList, item.ID)
	}
	return idList, nil
}
func (h *UpsertHandler) Check() error {
	if h.req.Material == nil {
		return errors.New("invalid material")
	}
	material := h.req.Material
	for _, item := range material {
		if item.UserId == 0 {
			return errors.New("invalid user_id")
		}
		if item.MaterialKey == "" {
			return errors.New("invalid material_key")
		}
		if item.MaterialValue == "" {
			return errors.New("invalid material_value")
		}
		if item.MaterialType == 0 {
			return errors.New("invalid material_type")
		}
	}

	return nil
}
