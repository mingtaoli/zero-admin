package logic

import (
	"context"
	"encoding/json"
	"go-zero-admin/api/internal/common/errorx"
	"go-zero-admin/rpc/pms/pmsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductCategoryUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCategoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductCategoryUpdateLogic {
	return ProductCategoryUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCategoryUpdateLogic) ProductCategoryUpdate(req types.UpdateProductCategoryReq) (*types.UpdateProductCategoryResp, error) {
	_, err := l.svcCtx.Pms.ProductCategoryUpdate(l.ctx, &pmsclient.ProductCategoryUpdateReq{
		Id:           req.Id,
		ParentId:     req.ParentId,
		Name:         req.Name,
		Level:        req.Level,
		ProductCount: req.ProductCount,
		ProductUnit:  req.ProductUnit,
		NavStatus:    req.NavStatus,
		ShowStatus:   req.ShowStatus,
		Sort:         req.Sort,
		Icon:         req.Icon,
		Keywords:     req.Keywords,
		Description:  req.Description,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.Errorf("更新商品分类参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新商品分类失败")
	}

	return &types.UpdateProductCategoryResp{
		Code:    "000000",
		Message: "更新商品分类成功",
	}, nil
}
