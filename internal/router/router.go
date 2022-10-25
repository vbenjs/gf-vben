package router

import (
	"Gf-Vben/internal/controller"
	"Gf-Vben/internal/middleware"
	"Gf-Vben/util"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 你可以将路由注册放到一个文件中管理，
// 也可以按照模块拆分到不同的文件中管理，
// 但统一都放到router目录下。
func init() {
	s := g.Server()

	s.Use(util.ResponseHandler, middleware.CORS)
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Map(g.Map{
			"login":    controller.User.Login,
			"register": controller.User.Register,
		})
		//group.Middleware(middleware.Auth)

	})
	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)
		group.Group("/user", func(group *ghttp.RouterGroup) {
			group.Bind(
				controller.User,
			)
		})
		group.Middleware(middleware.Casbin)

		group.Bind(
			controller.Curd,
		)

	})

}
