package logic

import (
	"context"
	"zero-basics/services/order/api/internal/svc"
	"zero-basics/services/order/api/internal/types"
	"zero-basics/services/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateRequest) (resp *types.UpdateResponse, err error) {
	_, err = l.svcCtx.OrderRpc.Update(l.ctx, &order.UpdateRequest{
		Id:     req.Id,
		Uid:    req.Uid,
		Pid:    req.Pid,
		Amount: req.Amount,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdateResponse{}, nil

	// newOrder := model.Order{
	// 	Id:     uint64(req.Id),
	// 	Uid:    uint64(req.Uid),
	// 	Pid:    uint64(req.Pid),
	// 	Amount: uint64(req.Amount),
	// 	Status: uint64(req.Status),
	// }
	// err = l.svcCtx.OrderModel.Update(l.ctx, &newOrder)
	// if err != nil {
	// 	return nil, err
	// }
	//
	// return &types.UpdateResponse{}, nil

}
