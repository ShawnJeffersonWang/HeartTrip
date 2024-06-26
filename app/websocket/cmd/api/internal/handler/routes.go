// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	websock "golodge/app/websocket/cmd/api/internal/handler/websock"
	"golodge/app/websocket/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// chat
				Method:  http.MethodGet,
				Path:    "/chat",
				Handler: websock.ChatHandler(serverCtx),
			},
			{
				// get inbox messages
				Method:  http.MethodGet,
				Path:    "/getInbox",
				Handler: websock.GetInboxHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/websocket/v1"),
	)
}
