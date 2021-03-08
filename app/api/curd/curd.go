package curd

import (
	"Gf-Vben/app/service/curd"
	"Gf-Vben/app/service/user"
	"Gf-Vben/library/response"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct {
}

func (c *Controller) Curd(r *ghttp.Request) {
	var cu curd.Curd
	switch r.GetString("interface") {
	case "user":
		req := new(user.Req)
		cu = req
	default:
		response.JsonExit(r, 1, "接口参数错误")
	}
	if err := r.Parse(cu); err != nil {
		response.JsonExit(r, 2, err.Error())
	}
	switch r.GetString("action") {
	case "list":
		result, err := cu.List()
		if err != nil {
			response.JsonExit(r, 3, err.Error())
		}
		response.JsonExit(r, 2, "", result)
	case "add":
		if err := cu.Add(); err != nil {
			response.JsonExit(r, 3, err.Error())
		}
		response.JsonExit(r, 2, "新增成功")
	case "edit":
		if err := cu.Edit(); err != nil {
			response.JsonExit(r, 3, err.Error())
		}
		response.JsonExit(r, 2, "修改成功")
	case "del":
		if err := cu.Del(); err != nil {
			response.JsonExit(r, 3, err.Error())
		}
		response.JsonExit(r, 2, "删除成功")
	default:
		response.JsonExit(r, 3, "接口参数错误")
	}
}
