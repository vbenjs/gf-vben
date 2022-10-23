package user

import (
	"Gf-Vben/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta `path:"/login" method:"post" summary:"登录请求" tags:"登录注册相关"`
}

type RegisterReq struct {
	g.Meta `path:"/register" method:"post" summary:"注册请求" tags:"登录注册相关"`
	model.RegisterReq
}
type InfoReq struct {
	g.Meta `path:"/info" method:"get" summary:"通过Token获取用户信息" `
}

type MenuReq struct {
	g.Meta `path:"/menu" method:"get" summary:"获取用户菜单"`
	Uid    int `p:"uid"`
	//user.MenuReq
}
