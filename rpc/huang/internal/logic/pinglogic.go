package logic

import (
	"context"
	"zero-basics/rpc/huang/huang"
	"zero-basics/rpc/huang/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *huang.Request) (*huang.Response, error) {
	// todo: add your logic here and delete this line

	return &huang.Response{
		Pong: "pong",
	}, nil
}
