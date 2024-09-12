package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero-basics/services/order/api/internal/config"
	"zero-basics/services/order/model"
)

type ServiceContext struct {
	Config     config.Config
	OrderModel model.OrderModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	//
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:     c,
		OrderModel: model.NewOrderModel(conn, c.CacheRedis),
	}
}
