package router

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

type Req struct {
	Id int `p:"id"`
}

type Router struct {
	Path      string `orm:"path" json:"path"`
	Name      string `orm:"name" json:"name"`
	Component string `orm:"component" json:"component"`
	Meta      `json:"meta"`
	Children  []*Router  `json:"children"`
	Status    int        `orm:"status" json:"status"`
	CreateAt  gtime.Time `orm:"create_at" json:"create_at"`
	OrderNo   int        `orm:"order_no" json:"order_no"`
}

type Meta struct {
	Title string `orm:"title" json:"title"`
	Icon  string `orm:"icon" json:"icon"`
}

func (r *Req) List() ([]*Router, error) {
	all, err := g.DB().Model("router").Where("status", 1).Order("parent").All()
	if err != nil {
		return nil, err

	}

	return getRouterList(all), nil
}

func getRouterList(data gdb.Result) []*Router {
	res := map[int]*Router{}
	for _, record := range data {
		var router *Router
		record.Struct(&router)

		if gconv.Int(record.Map()["parent"]) == 0 {
			router.Children = []*Router{}
			res[gconv.Int(record.Map()["id"])] = router
			continue
		}
		v := res[gconv.Int(record.Map()["parent"])]
		v.Children = append(v.Children, router)
		//g.Dump(v)
	}
	//TODO 路由排序
	var result []*Router
	for _, v := range res {
		result = append(result, v)
	}
	return result
}
