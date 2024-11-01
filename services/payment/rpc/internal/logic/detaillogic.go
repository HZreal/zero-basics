package logic

import (
	"context"
	"errors"
	"google.golang.org/grpc/status"
	"zero-basics/services/payment/model"

	"zero-basics/services/payment/rpc/internal/svc"
	"zero-basics/services/payment/rpc/types/pay"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *pay.DetailRequest) (*pay.DetailResponse, error) {
	// 查询支付是否存在
	res, err := l.svcCtx.PayModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(100, "支付不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	return &pay.DetailResponse{
		Id:     int64(res.Id),
		Uid:    int64(res.Uid),
		Oid:    int64(res.Oid),
		Amount: int64(res.Amount),
		Source: int64(res.Source),
		Status: int64(res.Status),
	}, nil

}
