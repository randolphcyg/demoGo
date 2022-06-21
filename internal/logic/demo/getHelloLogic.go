package demo

import (
	"context"

	"demoGo/internal/svc"
	"demoGo/internal/types"
	"demoGo/javastub/demopb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHelloLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHelloLogic {
	return &GetHelloLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHelloLogic) GetHello(req *types.HelloReq) (resp *types.HelloResp, err error) {
	logx.Info("http请求：", req.Name)
	reply, err := l.svcCtx.DemoRpcClient.SayHello(l.ctx, &demopb.HelloReq{Name: req.Name})
	if err != nil {
		return nil, err
	}

	logx.Info("返回消息:", reply)

	return &types.HelloResp{Message: "请求成功"}, nil
}
