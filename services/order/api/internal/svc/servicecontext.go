package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"zero-basics/services/order/api/internal/config"
	"zero-basics/services/order/model"
	"zero-basics/services/order/rpc/orderclient"
)

type ServiceContext struct {
	Config     config.Config
	OrderModel model.OrderModel

	OrderRpc orderclient.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	//
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:     c,
		OrderModel: model.NewOrderModel(conn, c.CacheRedis),
		OrderRpc:   orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
