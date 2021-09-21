package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

type Router struct {
	Path      string `orm:"path" json:"path"`
	Name      string `orm:"name" json:"name"`
	Component string `orm:"component" json:"component"`
	Meta      `json:"meta"`
	Children  []*Router  `json:"children"`
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

func (r *Req) Tree() (g.Map, error) {
	//user, err := casbin.CE.GetRolesForUser("vben", "router")
	//forUser := casbin.CE.GetPermissionsForUserInDomain("vben", "router")
	//if err != nil {
	//	return nil, err
	//
	//}
	var routers []*Router
	if err := g.DB().Model("router").Where("status", 1).Order("parent").Scan(&routers); err != nil {
		return nil, err
	}
	res := map[int]*Router{}

	for _, router := range routers {

		router.Children = make([]*Router, 0)
		res[router.Id] = router
		if r, ok := res[router.Parent]; ok {
			if len(r.Children) > 0 {
				if r.Children[0].OrderNo > router.OrderNo {
					r.Children = append([]*Router{router}, r.Children...)
					continue
				}
			}
			r.Children = append(r.Children, router)
		}

	}
	var result []*Router
	for _, v := range res {
		if v.Parent == 0 {
			if len(result) > 0 {
				if result[0].OrderNo > v.OrderNo {
					result = append([]*Router{v}, result...)
					continue
				}
			}
			result = append(result, v)
		}
	}

	return g.Map{"router": result}, nil
}
