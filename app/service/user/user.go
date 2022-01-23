package user

import (
	"Gf-Vben/app/model/entity"
	"Gf-Vben/app/service/internal/dao"
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
)

type LoginReq struct {
	Username string `json:"username" dc:"账号" v:"required"`
	Password string `json:"password" dc:"密码" v:"required"`
	Ctx      context.Context
}

func (r *LoginReq) Login() error {
	var u entity.User
	if err := dao.User.Ctx(r.Ctx).Where("username", r.Username).Scan(&u); err != nil {
		return err
	}

	if u.Id == 0 {
		return gerror.New("用户名或者密码不正确")
	}
	if u.Status == 0 {
		return gerror.New("用户已被禁用")
	}
	password, err := gmd5.Encrypt(r.Password)
	if err != nil {
		return err
	}

	if password != u.Password {
		return gerror.New("用户名或者密码不正确")
	}

	return nil
}

type RegisterReq struct {
	Ctx      context.Context
	Username string `p:"username" v:"required"`
	Pw       string `p:"password" v:"required"`
	Pw2      string `p:"password2" v:"required"`
}

func (r *RegisterReq) Register() error {
	if r.Pw != r.Pw2 {
		return gerror.New("密码不一致")
	}
	result, err := dao.User.Ctx(r.Ctx).One("username", r.Username)
	if err != nil {
		return err
	}
	if !result.IsEmpty() {
		return gerror.New("账号已存在")
	}

	pw, err := gmd5.Encrypt(r.Pw)
	if err != nil {
		return err
	}
	u := entity.User{
		Username: r.Username,
		Password: pw,
		Status:   1,
	}
	if _, err := dao.User.Ctx(r.Ctx).Insert(u); err != nil {
		return err
	}
	return nil
}
