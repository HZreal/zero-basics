package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"zero-basics/services/order/rpc/orderclient"
	"zero-basics/services/payment/model"
	"zero-basics/services/payment/rpc/internal/config"
	"zero-basics/services/user/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config

	PayModel model.PayModel

	UserRpc  userclient.User
	OrderRpc orderclient.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:   c,
		PayModel: model.NewPayModel(conn, c.CacheRedis),
		UserRpc:  userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		OrderRpc: orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
