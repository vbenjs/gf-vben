package router

import (
	"Gf-Vben/app/service/response"
	"Gf-Vben/app/service/router"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct {
}

func (c *Controller) List(r *ghttp.Request) {
	var req router.ListReq
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	list, err := req.List()
	if err != nil {
		response.JsonExit(r, 2, err.Error())

	}

	response.JsonExit(r, 0, "", list)

}
