// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	comment "github.com/deltafi-trade/looopx-api/internal/handler/comment"
	feeds "github.com/deltafi-trade/looopx-api/internal/handler/feeds"
	like "github.com/deltafi-trade/looopx-api/internal/handler/like"
	user "github.com/deltafi-trade/looopx-api/internal/handler/user"
	"github.com/deltafi-trade/looopx-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/user/login",
					Handler: user.LoginHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/captcha",
					Handler: user.CaptchaHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/logout",
					Handler: user.LogoutHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/:walletAddress",
					Handler: user.AccountHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/me",
					Handler: user.MeHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/settings",
					Handler: user.SettingsHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/avatar",
					Handler: user.AvatarHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/:user_id/following",
					Handler: user.FollowingHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/:user_id/follower",
					Handler: user.FollowerHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/notification",
					Handler: user.NotificationHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/user/like/:id/:type",
					Handler: like.ReadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/like/:id/:type",
					Handler: like.UpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/user/like/:id/:type",
					Handler: like.DeleteHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/feeds/list",
					Handler: feeds.ListHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/comment/:topic_id/:topic_type",
					Handler: comment.ReadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/comment/:topic_id/:topic_type",
					Handler: comment.UpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/comment/:comment_id",
					Handler: comment.DeleteHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/search",
					Handler: feeds.SearchHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1"),
	)
}