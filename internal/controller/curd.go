package controller

import (
	"Gf-Vben/api/v1"
	"Gf-Vben/internal/model"
	"Gf-Vben/internal/service"
	"Gf-Vben/util"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Curd = cCurd{}
)

type cCurd struct {
}

func (cCurd) Curd(ctx context.Context, req *v1.CurdReq) (res *util.JsonRes, err error) {
	r := g.RequestFromCtx(ctx)
	in := new(model.CurdReq)
	if err := r.Parse(&in); err != nil {
		return nil, err
	}
	return service.Curd().Curd(ctx, in)
}
