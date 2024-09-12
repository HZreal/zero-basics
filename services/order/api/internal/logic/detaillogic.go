package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"zero-basics/services/order/api/internal/svc"
	"zero-basics/services/order/api/internal/types"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.DetailRequest) (resp *types.DetailResponse, err error) {

	// res, err := l.svcCtx.OrderRpc.Detail(l.ctx, &order.DetailRequest{
	// 	Id: req.Id,
	// })
	// if err != nil {
	// 	return nil, err
	// }
	//
	// return &types.DetailResponse{
	// 	Id:     res.Id,
	// 	Uid:    res.Uid,
	// 	Pid:    res.Pid,
	// 	Amount: res.Amount,
	// 	Status: res.Status,
	// }, nil

	res, err := l.svcCtx.OrderModel.FindOne(l.ctx, uint64(req.Id))
	if err != nil {
		return nil, err
	}

	return &types.DetailResponse{
		Id:     int64(res.Id),
		Uid:    int64(res.Uid),
		Pid:    int64(res.Pid),
		Amount: int64(res.Amount),
		Status: int64(res.Status),
	}, nil

}
