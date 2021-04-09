package user

import (
	"Gf-Vben/app/service/response"
	"Gf-Vben/app/service/user"
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
func (c *Controller) Info(r *ghttp.Request) {
	var req *user.Req
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	res, err := req.List()
	if err != nil {
		response.JsonExit(r, 2, err.Error())
	}
	response.JsonExit(r, 0, "", res)
}
