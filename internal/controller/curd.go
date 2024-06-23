package controller

import (
	"Gf-Vben/internal/logic/router"
	"Gf-Vben/internal/logic/user"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/jinmao88/gf-utility/curd"
	"github.com/jinmao88/gf-utility/response"
)

var (
	Curd = cCurd{}
)

type cCurd struct {
}

func (cCurd) Curd(ctx context.Context, req *curd.CurdReq) (res *response.JsonRes, err error) {
	return curd.Controller(ctx, req, func(i string) (curd.Curd, error) {
		var cu curd.Curd
		switch i {
		case "user":
			cu = new(user.Req)
		case "router":
			cu = new(router.Req)

		default:
			return nil, gerror.NewCode(response.Code(1), "接口参数错误")
		}
		return cu, nil
	})
}
