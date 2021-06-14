package router

import (
	"Gf-Vben/app/api/curd"
	"Gf-Vben/app/api/user"
	"Gf-Vben/app/service/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 你可以将路由注册放到一个文件中管理，
// 也可以按照模块拆分到不同的文件中管理，
// 但统一都放到router目录下。
func init() {
	s := g.Server()
	s.BindHandler("/*", func(r *ghttp.Request) {

	})
	s.BindMiddleware("/*", middleware.CORS)
	s.BindHandler("POST:/login", middleware.GfJWTMiddleware.LoginHandler)
	s.BindHandler("POST:/register", user.Register)
	// 分组路由注册方式
	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)
		group.ALL("/user/{.method}", new(user.Controller))
		group.Middleware(middleware.Casbin)
		group.ALL("/curd", new(curd.Controller).Curd)

	})
}
