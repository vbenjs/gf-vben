package user

import (
	"Gf-Vben/internal/dao"
	"Gf-Vben/internal/model"
	"Gf-Vben/internal/model/entity"
	"Gf-Vben/internal/service"
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/jinmao88/gf-utility/menu"
)

type sUser struct {
}

func init() {
	service.RegisterUser(New())
}
func New() *sUser {
	return &sUser{}
}

func (s *sUser) Register(ctx context.Context, in model.RegisterReq) error {

	if in.Pw != in.Pw2 {
		return gerror.New("密码不一致")
	}
	result, err := dao.User.Ctx(ctx).One("username", in.Username)
	if err != nil {
		return err
	}
	if !result.IsEmpty() {
		return gerror.New("账号已存在")
	}

	pw, err := gmd5.Encrypt(in.Pw)
	if err != nil {
		return err
	}
	u := entity.User{
		Username: in.Username,
		Password: pw,
		Status:   1,
	}
	if _, err := dao.User.Ctx(ctx).Insert(u); err != nil {
		return err
	}
	return nil
}

//type MenuReq struct {
//	Uid string `p:"uid"`
//	Ctx context.Context
//}

func (s *sUser) Menu(ctx context.Context) ([]*menu.Menu, error) {
	//casbin.CE.LoadPolicy()
	//var p []string
	//permissions := casbin.CE.GetPermissionsForUserInDomain(r.Uid, "menu")
	//for _, permission := range permissions {
	//	p = append(p, permission[2])
	//}
	var routers []*menu.Menu
	if err := dao.Router.Ctx(ctx).Where(dao.Router.Columns().Status, 1).Order(dao.Router.Columns().Parent).Scan(&routers); err != nil {
		return nil, err
	}
	return menu.BuildRouter(routers), nil
}

func (s *sUser) Info(ctx context.Context, uid int) (gdb.Record, error) {

	one, err := dao.User.Ctx(ctx).Where(dao.User.Columns().Id, uid).FieldsEx(dao.User.Columns().Password).One()
	if err != nil {
		return nil, err
	}
	if one.IsEmpty() {
		return nil, gerror.New("用户不存在")
	}

	return one, nil
}
