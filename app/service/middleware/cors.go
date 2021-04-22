package middleware

import "github.com/gogf/gf/net/ghttp"

// 允许接口跨域请求
func CORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	corsOptions.AllowHeaders += "Access-Token"
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}
