// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2

package handler

import (
	"net/http"

	"zero-basics/services/product/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/product/create",
				Handler: CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/product/detail",
				Handler: DetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/product/remove",
				Handler: RemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/product/update",
				Handler: UpdateHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}