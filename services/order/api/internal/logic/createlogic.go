package logic

import (
	"context"
	"zero-basics/services/order/api/internal/svc"
	"zero-basics/services/order/api/internal/types"
	"zero-basics/services/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	res, err := l.svcCtx.OrderRpc.Create(l.ctx, &order.CreateRequest{
		Uid:    req.Uid,
		Pid:    req.Pid,
		Amount: req.Amount,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateResponse{
		Id: res.Id,
	}, nil

	// newOrder := model.Order{
	// 	Uid:    uint64(req.Uid),
	// 	Pid:    uint64(req.Pid),
	// 	Amount: uint64(req.Amount),
	// 	Status: 0,
	// }
	// res, err := l.svcCtx.OrderModel.Insert(l.ctx, &newOrder)
	// if err != nil {
	// 	return nil, err
	// }
	//
	// var insertUserId int64
	// insertUserId, err = res.LastInsertId()
	// if err != nil {
	// 	return nil, status.Error(500, err.Error())
	// }
	//
	// return &types.CreateResponse{
	// 	Id: insertUserId,
	// }, nil

}
