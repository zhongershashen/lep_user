package handler

import (
	"bupt/lep_user/biz/logic/material"
	"bupt/lep_user/kitex_gen/base"
	"bupt/lep_user/kitex_gen/lep_user"
	"context"
)

func MaterialList(ctx context.Context, req *lep_user.MaterialListRequest) (*lep_user.MaterialListResp, error) {
	resp := new(lep_user.MaterialListResp)
	resp.BaseResp = base.NewBaseResp()
	list, total, err := material.NewMaterialListHandler(ctx, req).GetMaterialList()
	if err != nil {
		return nil, err
	}
	resp.MaterialList = list
	resp.Total = total
	return resp, err
}
func UpsertMaterial(ctx context.Context, req *lep_user.UpsertMaterialRequest) (*lep_user.UpsertMaterialResp, error) {
	resp := new(lep_user.UpsertMaterialResp)
	resp.BaseResp = base.NewBaseResp()
	handler := material.NewUpsertHandler(ctx, req)
	err := handler.Check()
	if err != nil {
		return nil, err
	}
	id, err := handler.UpsertMaterial()
	if err != nil {
		return nil, err
	}
	resp.MaterialId = id
	return resp, err
}
func DeleteMaterial(ctx context.Context, req *lep_user.DeleteMaterialRequest) (*lep_user.DeleteMaterialResp, error) {
	resp := new(lep_user.DeleteMaterialResp)
	resp.BaseResp = base.NewBaseResp()
	handler := material.NewDeleteHandler(ctx, req)
	err := handler.Delete()
	if err != nil {
		return nil, err
	}
	return resp, err
}
