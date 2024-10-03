package logic

import (
	"context"
	"zero-basics/services/order/api/internal/svc"
	"zero-basics/services/order/api/internal/types"
	"zero-basics/services/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.ListRequest) (resp []*types.ListResponse, err error) {
	res, err := l.svcCtx.OrderRpc.List(l.ctx, &order.ListRequest{
		Uid: req.Uid,
	})
	if err != nil {
		return nil, err
	}

	orderList := make([]*types.ListResponse, 0)
	for _, item := range res.Data {
		orderList = append(orderList, &types.ListResponse{
			Id:     item.Id,
			Uid:    item.Uid,
			Pid:    item.Pid,
			Amount: item.Amount,
			Status: item.Status,
		})
	}

	return orderList, nil

}
