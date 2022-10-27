package controller

import (
	"Gf-Vben/api/v1/user"
	"Gf-Vben/internal/middleware"
	"Gf-Vben/internal/service"
	"Gf-Vben/util"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	User = cUser{}
)

type cUser struct {
}

func (cUser) Login(ctx context.Context, req *user.LoginReq) (res *util.JsonRes, err error) {
	res = new(util.JsonRes)
	token, _ := middleware.GfJWTMiddleware.LoginHandler(ctx)
	res.Data = g.Map{"token": token}
	return
}

func (cUser) Register(ctx context.Context, req *user.RegisterReq) (res *util.JsonRes, err error) {
	res = new(util.JsonRes)
	err = service.User().Register(ctx, req.RegisterReq)
	res.Message = "注册成功"
	return
}

func (cUser) Info(ctx context.Context, req *user.InfoReq) (res *util.JsonRes, err error) {
	res = new(util.JsonRes)
	res.Data, err = service.User().Info(ctx, req.Uid)
	return
}

func (cUser) Menu(ctx context.Context, req *user.MenuReq) (res *util.JsonRes, err error) {
	res = new(util.JsonRes)
	res.Data, err = service.User().Menu(ctx)

	return
}
