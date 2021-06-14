package middleware

import (
	"Gf-Vben/app/service/response"
	"github.com/casbin/casbin/v2"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

var CE *casbin.Enforcer

func Casbin(r *ghttp.Request) {

	if ok, _ := CE.Enforce(r.GetParam("role"), r.GetString("interface"), r.GetString("action")); ok {
		r.Middleware.Next()
	} else {
		glog.Println(r.GetParam("role"), r.GetString("interface"), r.GetString("action"))
		response.JsonExit(r, 1, "无此权限")
	}

}
