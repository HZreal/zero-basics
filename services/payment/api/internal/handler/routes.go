// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2

package handler

import (
	"net/http"

	"zero-basics/services/payment/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/pay/callback",
				Handler: CallbackHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/pay/create",
				Handler: CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/pay/detail",
				Handler: DetailHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
