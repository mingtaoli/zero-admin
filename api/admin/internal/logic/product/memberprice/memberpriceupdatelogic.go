package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberPriceUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberPriceUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberPriceUpdateLogic {
	return MemberPriceUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberPriceUpdateLogic) MemberPriceUpdate(req types.UpdateMemberPriceReq) (*types.UpdateMemberPriceResp, error) {
	_, err := l.svcCtx.MemberPriceService.MemberPriceUpdate(l.ctx, &pmsclient.MemberPriceUpdateReq{
		Id:              req.Id,
		ProductId:       req.ProductId,
		MemberLevelId:   req.MemberLevelId,
		MemberPrice:     int64(req.MemberPrice),
		MemberLevelName: req.MemberLevelName,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新会员价格信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新会员价格失败")
	}

	return &types.UpdateMemberPriceResp{
		Code:    "000000",
		Message: "更新会员价格成功",
	}, nil
}
