package logic

import (
	"context"
	"errors"
	"google.golang.org/grpc/status"
	"zero-basics/services/order/rpc/types/order"
	"zero-basics/services/payment/model"
	"zero-basics/services/user/rpc/types/user"

	"zero-basics/services/payment/rpc/internal/svc"
	"zero-basics/services/payment/rpc/types/pay"

	"github.com/zeromicro/go-zero/core/logx"
)

type CallbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallbackLogic {
	return &CallbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CallbackLogic) Callback(in *pay.CallbackRequest) (*pay.CallbackResponse, error) {
	// 查询用户是否存在
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: in.Uid,
	})
	if err != nil {
		return nil, err
	}

	// 查询订单是否存在
	_, err = l.svcCtx.OrderRpc.Detail(l.ctx, &order.DetailRequest{
		Id: in.Oid,
	})
	if err != nil {
		return nil, err
	}

	// 查询支付是否存在
	res, err := l.svcCtx.PayModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(100, "支付不存在")
		}
		return nil, status.Error(500, err.Error())
	}
	// 支付金额与订单金额不符
	if uint64(in.Amount) != res.Amount {
		return nil, status.Error(100, "支付金额与订单金额不符")
	}

	res.Source = uint64(in.Source)
	res.Status = uint64(in.Status)

	err = l.svcCtx.PayModel.Update(l.ctx, res)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// 更新订单支付状态
	_, err = l.svcCtx.OrderRpc.Paid(l.ctx, &order.PaidRequest{
		Id: in.Oid,
	})
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &pay.CallbackResponse{}, nil

}
