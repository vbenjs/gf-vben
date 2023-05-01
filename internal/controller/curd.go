package controller

import (
	"Gf-Vben/api/v1"
	"Gf-Vben/internal/service"
	"context"
	"github.com/jinmao88/gf-utility/response"
)

var (
	Curd = cCurd{}
)

type cCurd struct {
}

func (cCurd) Curd(ctx context.Context, req *v1.CurdReq) (res *response.JsonRes, err error) {
	return service.Curd().Curd(ctx, &req.CurdReq)
}
