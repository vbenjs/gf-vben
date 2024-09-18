package router

import (
	"Gf-Vben/internal/controller"
	"Gf-Vben/internal/middleware"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
)

// 你可以将路由注册放到一个文件中管理，
// 也可以按照模块拆分到不同的文件中管理，
// 但统一都放到router目录下。
func init() {
	s := g.Server()

	s.Use(middleware.ResponseHandler, middleware.CORS)
	middleware.GfTokenInstance = &gtoken.GfToken{
		LoginPath:        "/auth/login",
		LoginBeforeFunc:  middleware.Login,
		LogoutPath:       "/auth/logout",
		MultiLogin:       true,
		AuthExcludePaths: g.SliceStr{"/api/auth/register"},
		LoginAfterFunc:   middleware.LoginAfterFunc,
		AuthAfterFunc:    middleware.AuthAfterFunc,
	}

	s.Group("/api", func(group *ghttp.RouterGroup) {
		middleware.GfTokenInstance.Middleware(gctx.New(), group)

		group.Group("/auth", func(rg *ghttp.RouterGroup) {
			rg.Bind(controller.User)
		})

		group.Group("/menu", func(rg *ghttp.RouterGroup) {
			rg.Bind(controller.Menu)
		})

		group.Bind(
			controller.Curd,
		)

	})

}
