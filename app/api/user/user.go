package user

import (
	"Gf-Vben/app/service/user"
	"Gf-Vben/library/response"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct {
}

func Register(r *ghttp.Request) {
	var req *user.RegisterReq
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := req.Register(); err != nil {
		response.JsonExit(r, 2, err.Error())
	}
	response.JsonExit(r, 0, "注册成功")

}
