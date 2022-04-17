package middleware

import (
	"Gf-Vben/app/service/casbin"
	"Gf-Vben/app/util"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
)

func Casbin(r *ghttp.Request) {
	var req casbin.Req
	if err := r.Parse(&req); err != nil {
		util.JsonExit(r, 1, "权限失效")
	}
	if gstr.Contains(r.RequestURI, "curd") {
		req.Domain = "curd"
	}
	if err := req.Check(); err != nil {
		util.JsonExit(r, 2, err.Error())
	}
	r.Middleware.Next()

}
