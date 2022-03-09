package user

import (
	"Gf-Vben/app/service/middleware"
	"Gf-Vben/app/service/user"
	"Gf-Vben/app/util"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type Api struct {
}
type LoginReq struct {
	g.Meta `path:"/login" method:"post" summary:"登录请求" tags:"登录注册相关"`
}

//type RefreshTokenReq struct {
//	g.Meta `path:"/refresh_token" method:"post"`
//}
//
//type AuthLogoutReq struct {
//	g.Meta `path:"/logout" method:"post"`
//}
//

func (Api) Login(ctx context.Context, req *LoginReq) (res *util.JsonRes, err error) {
	res = new(util.JsonRes)
	token, _ := middleware.GfJWTMiddleware.LoginHandler(ctx)
	res.Data = g.Map{"token": token}
	return
}

//func (Api) RefreshToken(ctx context.Context, req *api.AuthRefreshTokenReq) (res *api.AuthRefreshTokenRes, err error) {
//	res = &api.AuthRefreshTokenRes{}
//	res.Token, res.Expire = service.Auth().RefreshHandler(ctx)
//	return
//}
//
//func (Api) Logout(ctx context.Context, req *api.AuthLogoutReq) (res *api.AuthLogoutRes, err error) {
//	service.Auth().LogoutHandler(ctx)
//	return
//}

//func (Api) Login(ctx context.Context, req *LoginReq) (res *util.JsonRes, err error) {
//	token := req.Username + req.Password
//	g.Dump(token)
//	res = new(util.JsonRes)
//
//	if err := req.LoginReq.Login(); err != nil {
//		return res, gerror.WrapCode(util.Code(1), err)
//	}
//	res.Data = g.Map{"token": token}
//
//	return
//}

type RegisterReq struct {
	g.Meta `path:"/register" method:"post" summary:"注册请求" tags:"登录注册相关"`
	user.RegisterReq
}

func (Api) Register(ctx context.Context, req *RegisterReq) (res *util.JsonRes, err error) {

	res = new(util.JsonRes)
	if err := req.RegisterReq.Register(); err != nil {
		return nil, gerror.WrapCode(util.Code(1), err)
	}
	res.Message = "注册成功"

	return
}

type Api2 struct {
}
type InfoReq struct {
	g.Meta `path:"/info" method:"get" summary:"通过Token获取用户信息" `
}

func (Api2) Info(ctx context.Context, req *InfoReq) (res *util.JsonRes, err error) {

	res = new(util.JsonRes)
	res.Data = g.Map{
		"username": "vben",
		"roles":    []string{"admin"},
	}

	return
}
