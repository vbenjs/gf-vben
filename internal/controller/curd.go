package controller

import (
	"Gf-Vben/api/v1"
	"Gf-Vben/internal/service"
	"Gf-Vben/util"
	"context"
)

var (
	Curd = cCurd{}
)

type cCurd struct {
}

func (cCurd) Curd(ctx context.Context, req *v1.CurdReq) (res *util.JsonRes, err error) {
	return service.Curd().Curd(ctx, &req.CurdReq)
}
