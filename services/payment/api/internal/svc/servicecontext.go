package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zero-basics/services/payment/api/internal/config"
	"zero-basics/services/payment/rpc/payclient"
)

type ServiceContext struct {
	Config config.Config

	PayRpc payclient.Pay
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		PayRpc: payclient.NewPay(zrpc.MustNewClient(c.PayRpc)),
	}
}
