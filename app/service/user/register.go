package user

import (
	"Gf-Vben/app/dao"
	"Gf-Vben/app/model"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/errors/gerror"
)

type RegisterReq struct {
	Username string `p:"username" v:"required"`
	Pw       string `p:"password" v:"required"`
	Pw2      string `p:"password2" v:"required"`
}

func (r *RegisterReq) Register() error {
	result, err := dao.AppUser.FindOne("username", r.Username)
	if err != nil {
		return err
	}
	if result != nil {
		return gerror.New("账号已存在")
	}
	if r.Pw != r.Pw2 {
		return gerror.New("密码不一致")
	}
	pw, err := gmd5.Encrypt(r.Pw)
	if err != nil {
		return err
	}
	u := model.AppUser{
		Username: r.Username,
		Password: pw,
		Status:   1,
	}
	if _, err := dao.AppUser.Save(u); err != nil {
		return err
	}
	return nil
}
