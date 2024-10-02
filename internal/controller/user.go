package controller

import (
	"Gf-Vben/api/v1/user"
	"Gf-Vben/internal/middleware"
	"Gf-Vben/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/jinmao88/gf-utility/response"
)

var (
	User = cUser{}
)

type cUser struct {
}

func (cUser) Register(ctx context.Context, req *user.RegisterReq) (res response.JsonRes, err error) {
	r := new(middleware.Resp)
	err = service.User().Register(ctx, req.RegisterReq)
	r.Message = "注册成功"
	return r, err
}

func (cUser) Info(ctx context.Context, req *user.InfoReq) (res response.JsonRes, err error) {
	r := new(middleware.Resp)
	r.Data, err = service.User().Info(ctx, req.Uid)
	return r, err
}

func (cUser) Menu(ctx context.Context, req *user.MenuReq) (res response.JsonRes, err error) {
	r := new(middleware.Resp)
	r.Data, err = service.User().Menu(ctx)
	return r, err
}

func (cUser) AccessCodes(ctx context.Context, req *user.AccessCodeReq) (res response.JsonRes, err error) {
	r := new(middleware.Resp)
	codes, err := service.User().AccessCode(ctx, req.Role)
	r.Data = g.Map{
		"codes": codes,
		"uid":   req.Uid,
	}
	return r, err
}
