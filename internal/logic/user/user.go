package user

import (
	"Gf-Vben/internal/dao"
	"Gf-Vben/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"sort"
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
	//casbin.CE.LoadPolicy()
	//var p []string
	//permissions := casbin.CE.GetPermissionsForUserInDomain(r.Uid, "menu")
	//for _, permission := range permissions {
	//	p = append(p, permission[2])
	//}
	var routers []*Menu
	if err := dao.Router.Ctx(r.Ctx).Where(dao.Router.Columns().Status, 1).Order(dao.Router.Columns().Parent).Scan(&routers); err != nil {
		return nil, err
	}
	return BuildRouter(routers), nil
}

type Menu struct {
	Path      string `orm:"path" json:"path"`
	Name      string `orm:"name" json:"name"`
	Component string `orm:"component" json:"component"`
	Meta      `json:"meta"`
	Children  Menus      `json:"children"`
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

type Menus []*Menu

func (m Menus) Len() int {
	return len(m)
}

func (m Menus) Less(i, j int) bool {
	return m[i].OrderNo < m[j].OrderNo
}

func (m Menus) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func BuildRouter(routers []*Menu) (result Menus) {
	res := map[int]*Menu{}
	for _, router := range routers {
		router.Children = make(Menus, 0)
		res[router.Id] = router
		if router.Parent == 0 {
			result = append(result, router)
		}
		if r, ok := res[router.Parent]; ok {
			r.Children = append(r.Children, router)
		}
	}

	sort.Sort(result)
	for _, menu := range result {
		if len(menu.Children) > 0 {
			sort.Sort(menu.Children)
		}
	}
	return
}
