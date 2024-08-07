// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "golodge/app/usercenter/cmd/api/internal/handler/user"
	"golodge/app/usercenter/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// login
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				// register
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: user.RegisterHandler(serverCtx),
			},
		},
		rest.WithPrefix("/usercenter/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// upload
				Method:  http.MethodPost,
				Path:    "/upload",
				Handler: user.FileUploadHandler(serverCtx),
			},
			{
				// add wish list
				Method:  http.MethodPost,
				Path:    "/user/addWishList",
				Handler: user.AddWishListHandler(serverCtx),
			},
			{
				// browing history list
				Method:  http.MethodPost,
				Path:    "/user/browinghistory",
				Handler: user.HistoryListHandler(serverCtx),
			},
			{
				// clear history
				Method:  http.MethodPost,
				Path:    "/user/clearHistory",
				Handler: user.ClearHistoryHandler(serverCtx),
			},
			{
				// get user info
				Method:  http.MethodPost,
				Path:    "/user/detail",
				Handler: user.DetailHandler(serverCtx),
			},
			{
				// remove history
				Method:  http.MethodPost,
				Path:    "/user/removeHistory",
				Handler: user.RemoveHistoryHandler(serverCtx),
			},
			{
				// remove wish list
				Method:  http.MethodPost,
				Path:    "/user/removeWishList",
				Handler: user.RemoveWishListHandler(serverCtx),
			},
			{
				// update user info
				Method:  http.MethodPut,
				Path:    "/user/updateUserInfo",
				Handler: user.UpdateUserInfoHandler(serverCtx),
			},
			{
				// show user list
				Method:  http.MethodPost,
				Path:    "/user/userList",
				Handler: user.UserListHandler(serverCtx),
			},
			{
				// wishList
				Method:  http.MethodPost,
				Path:    "/user/wishList",
				Handler: user.WishListHandler(serverCtx),
			},
			{
				// wechat mini auth
				Method:  http.MethodPost,
				Path:    "/user/wxMiniAuth",
				Handler: user.WxMiniAuthHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/usercenter/v1"),
	)
}
