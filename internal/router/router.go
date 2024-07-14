package router

import (
	"Gf-Vben/internal/controller"
	"Gf-Vben/internal/middleware"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 你可以将路由注册放到一个文件中管理，
// 也可以按照模块拆分到不同的文件中管理，
// 但统一都放到router目录下。
func init() {
	s := g.Server()

	s.Use(middleware.ResponseHandler, middleware.CORS)

	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Group("/auth", func(rg *ghttp.RouterGroup) {
			rg.Map(g.Map{
				"login":    controller.User.Login,
				"register": controller.User.Register,
			})
		})
		group.Middleware(middleware.Auth)
		group.Group("/auth", func(rg *ghttp.RouterGroup) {
			rg.Map(g.Map{
				"menu":           controller.User.Menu,
				"getUserInfo":    controller.User.Info,
				"getAccessCodes": controller.User.AccessCodes,
			})
		})

		group.Bind(
			controller.Curd,
		)

	})

}
