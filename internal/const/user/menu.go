package user

import (
	"github.com/gogf/gf/v2/os/gtime"
	"sort"
)

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
		if router.Parent == 0 {
			result = append(result, router)
		}
		if r, ok := res[router.Parent]; ok {
			r.Children = append(r.Children, router)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].OrderNo < result[j].OrderNo
	})
	for _, menu := range result {
		if len(menu.Children) > 0 {
			sort.Slice(menu.Children, func(i, j int) bool {
				return menu.Children[i].OrderNo < menu.Children[j].OrderNo
			})

		}
	}
	return
}
