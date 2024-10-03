package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"zero-basics/services/product/model"

	"zero-basics/services/product/rpc/internal/svc"
	"zero-basics/services/product/rpc/types/product"

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

func (l *CreateLogic) Create(in *product.CreateRequest) (*product.CreateResponse, error) {
	newProduct := model.Product{
		Name:   in.Name,
		Desc:   in.Desc,
		Stock:  uint64(in.Stock),
		Amount: uint64(in.Amount),
		Status: uint64(in.Status),
	}

	res, err := l.svcCtx.ProductModel.Insert(l.ctx, &newProduct)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	insertId, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	newProduct.Id = uint64(insertId)

	return &product.CreateResponse{Id: insertId}, nil
}
