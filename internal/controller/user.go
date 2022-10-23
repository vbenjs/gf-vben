package controller

import (
	"Gf-Vben/api/v1/user"
	"Gf-Vben/internal/middleware"
	"Gf-Vben/internal/service"
	"Gf-Vben/util"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
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

	if err := service.User().Register(ctx, req.RegisterReq); err != nil {
		return nil, gerror.WrapCode(util.Code(1), err)
	}
	res.Message = "注册成功"
	return
}

func (cUser) Info(ctx context.Context, req *user.InfoReq) (res *util.JsonRes, err error) {

	res = new(util.JsonRes)
	res.Data = g.Map{
		"username": "vben",
		"roles":    []string{"admin"},
	}

	return
}

func (cUser) Menu(ctx context.Context, req *user.MenuReq) (res *util.JsonRes, err error) {
	g.Dump(req)
	res = new(util.JsonRes)
	menu, err := service.User().Menu(ctx)

	if err != nil {
		return nil, gerror.WrapCode(util.Code(1), err)
	}
	res.Data = menu
	return
}
