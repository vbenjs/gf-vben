package app

import (
	"Gf-Vben/app/api/curd"
	"Gf-Vben/app/api/user"
	"Gf-Vben/app/service/middleware"
	"Gf-Vben/app/util"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 你可以将路由注册放到一个文件中管理，
// 也可以按照模块拆分到不同的文件中管理，
// 但统一都放到router目录下。
func router() {
	s := g.Server()

	s.Use(util.ResponseHandler, middleware.CORS)
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Bind(
			new(user.Api),
		)

		group.Group("/user", func(group *ghttp.RouterGroup) {
			group.Bind(
				new(user.Api2),
			)
		})

	})
	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Bind(
			new(curd.Api),
		)
	})
	//s.BindHandler("/login", .Login)
	//s.BindHandler("POST:/register", user.Register)
	// 分组路由注册方式
	//s.Group("/api", func(group *ghttp.RouterGroup) {
	//middleware.Gtoken.Middleware(group)
	//group.Middleware(middleware.Auth)
	//group.Bind()
	//group.ALL("/user/info", user.Info)
	//group.Middleware(middleware.Casbin)

	//group.ALL("/curd", curd.Curd)

	//})
}
