package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"zero-basics/services/order/rpc/types/order"
	"zero-basics/services/payment/model"
	"zero-basics/services/user/rpc/types/user"

	"zero-basics/services/payment/rpc/internal/svc"
	"zero-basics/services/payment/rpc/types/pay"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *pay.CreateRequest) (*pay.CreateResponse, error) {
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

	// 查询订单是否已经创建支付
	_, err = l.svcCtx.PayModel.FindOneByOid(l.ctx, in.Oid)
	if err == nil {
		return nil, status.Error(100, "订单已创建支付")
	}

	newPay := model.Pay{
		Uid:    uint64(in.Uid),
		Oid:    uint64(in.Oid),
		Amount: uint64(in.Amount),
		Source: 0,
		Status: 0,
	}

	res, err := l.svcCtx.PayModel.Insert(l.ctx, &newPay)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	insertId, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	newPay.Id = uint64(insertId)

	return &pay.CreateResponse{
		Id: insertId,
	}, nil

}
