// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: order.proto

package server

import (
	"context"

	"zero-basics/services/order/rpc/internal/logic"
	"zero-basics/services/order/rpc/internal/svc"
	"zero-basics/services/order/rpc/types/order"
)

type OrderServer struct {
	svcCtx *svc.ServiceContext
	order.UnimplementedOrderServer
}

func NewOrderServer(svcCtx *svc.ServiceContext) *OrderServer {
	return &OrderServer{
		svcCtx: svcCtx,
	}
}

func (s *OrderServer) Create(ctx context.Context, in *order.CreateRequest) (*order.CreateResponse, error) {
	l := logic.NewCreateLogic(ctx, s.svcCtx)
	return l.Create(in)
}

func (s *OrderServer) Update(ctx context.Context, in *order.UpdateRequest) (*order.UpdateResponse, error) {
	l := logic.NewUpdateLogic(ctx, s.svcCtx)
	return l.Update(in)
}

func (s *OrderServer) Remove(ctx context.Context, in *order.RemoveRequest) (*order.RemoveResponse, error) {
	l := logic.NewRemoveLogic(ctx, s.svcCtx)
	return l.Remove(in)
}

func (s *OrderServer) Detail(ctx context.Context, in *order.DetailRequest) (*order.DetailResponse, error) {
	l := logic.NewDetailLogic(ctx, s.svcCtx)
	return l.Detail(in)
}

func (s *OrderServer) List(ctx context.Context, in *order.ListRequest) (*order.ListResponse, error) {
	l := logic.NewListLogic(ctx, s.svcCtx)
	return l.List(in)
}

func (s *OrderServer) Paid(ctx context.Context, in *order.PaidRequest) (*order.PaidResponse, error) {
	l := logic.NewPaidLogic(ctx, s.svcCtx)
	return l.Paid(in)
}
