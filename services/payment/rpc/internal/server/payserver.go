// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: pay.proto

package server

import (
	"context"

	"zero-basics/services/payment/rpc/internal/logic"
	"zero-basics/services/payment/rpc/internal/svc"
	"zero-basics/services/payment/rpc/types/pay"
)

type PayServer struct {
	svcCtx *svc.ServiceContext
	pay.UnimplementedPayServer
}

func NewPayServer(svcCtx *svc.ServiceContext) *PayServer {
	return &PayServer{
		svcCtx: svcCtx,
	}
}

func (s *PayServer) Create(ctx context.Context, in *pay.CreateRequest) (*pay.CreateResponse, error) {
	l := logic.NewCreateLogic(ctx, s.svcCtx)
	return l.Create(in)
}

func (s *PayServer) Detail(ctx context.Context, in *pay.DetailRequest) (*pay.DetailResponse, error) {
	l := logic.NewDetailLogic(ctx, s.svcCtx)
	return l.Detail(in)
}

func (s *PayServer) Callback(ctx context.Context, in *pay.CallbackRequest) (*pay.CallbackResponse, error) {
	l := logic.NewCallbackLogic(ctx, s.svcCtx)
	return l.Callback(in)
}