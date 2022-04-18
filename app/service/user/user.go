package user

import (
	"Gf-Vben/app/model/entity"
	"Gf-Vben/app/service/casbin"
	"Gf-Vben/app/service/internal/dao"
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
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

type MenuReq struct {
	Uid string `p:"uid"`
	Ctx context.Context
}

func (r *MenuReq) Menu() ([]*Menu, error) {
	casbin.CE.LoadPolicy()
	var p []string
	permissions := casbin.CE.GetPermissionsForUserInDomain(r.Uid, "menu")
	for _, permission := range permissions {
		p = append(p, permission[2])
	}
	var routers []*Menu
	if err := g.DB().Model("router").Where("status", 1).Where("permission", p).Order("parent").Scan(&routers); err != nil {
		return nil, err
	}
	return BuildRouter(routers), nil
}

type Menu struct {
	Path      string `orm:"path" json:"path"`
	Name      string `orm:"name" json:"name"`
	Component string `orm:"component" json:"component"`
	Meta      `json:"meta"`
	Children  []*Menu    `json:"children"`
	Status    int        `orm:"status" json:"status"`
	CreateAt  gtime.Time `orm:"create_at" json:"create_at"`
	OrderNo   int        `orm:"order_no" json:"order_no"`
	Id        int        `orm:"id" json:"id"`
	Parent    int        `orm:"parent" json:"parent"`
}

type Meta struct {
	Title string `orm:"title" json:"title"`
	Icon  string `orm:"icon" json:"icon"`
}

func BuildRouter(routers []*Menu) (result []*Menu) {
	res := map[int]*Menu{}
	for _, router := range routers {
		router.Children = make([]*Menu, 0)
		res[router.Id] = router
		if r, ok := res[router.Parent]; ok {
			if len(r.Children) > 0 {
				if r.Children[0].OrderNo > router.OrderNo {
					r.Children = append([]*Menu{router}, r.Children...)
					continue
				}
			}
			r.Children = append(r.Children, router)
		}

	}
	for _, v := range res {
		if v.Parent == 0 {
			if len(result) > 0 {
				if result[0].OrderNo > v.OrderNo {
					result = append([]*Menu{v}, result...)
					continue
				}
			}
			result = append(result, v)
		}
	}
	return
}
