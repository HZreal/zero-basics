package logic

import (
	"context"
	"zero-basics/services/order/api/internal/svc"
	"zero-basics/services/order/api/internal/types"
	"zero-basics/services/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveLogic) Remove(req *types.RemoveRequest) (resp *types.RemoveResponse, err error) {
	_, err = l.svcCtx.OrderRpc.Remove(l.ctx, &order.RemoveRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.RemoveResponse{}, nil

	// err = l.svcCtx.OrderModel.Delete(l.ctx, uint64(req.Id))
	// if err != nil {
	// 	return nil, err
	// }
	//
	// return &types.RemoveResponse{}, nil

}
