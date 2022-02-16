package logic

import (
	"context"

	"go-zero-admin/rpc/cms/internal/svc"
	"go-zero-admin/rpc/cms/types/cms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GreetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGreetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GreetLogic {
	return &GreetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GreetLogic) Greet(in *cms.StreamReq) (*cms.StreamResp, error) {
	// todo: add your logic here and delete this line

	return &cms.StreamResp{}, nil
}