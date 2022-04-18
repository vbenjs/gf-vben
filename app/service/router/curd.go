package router

import (
	"Gf-Vben/app/model/entity"
	"Gf-Vben/app/service/internal/dao"
	"context"
	tree "github.com/azhengyongqin/golang-tree-menu"
	"github.com/gogf/gf/v2/frame/g"
)

type Req struct {
	Page     int `p:"page"`
	PageSize int `p:"page_size"`
	Ctx      context.Context
	Query
}

type Query struct {
	Id        int    `p:"id"`
	Uid       int    `p:"uid"`
	Name      string `p:"name"`
	Parent    string `p:"parent"`
	Title     string `p:"title"`
	OrderNo   int    `p:"order_no"`
	Component string `p:"component"`
	Icon      string `p:"icon"`
	Redirect  string `p:"redirect"`
	Path      string `p:"path"`
	Status    int    `p:"status"`
}

func (r *Req) List() (g.Map, error) {
	all, err := g.DB().Model("router").Order("parent,orderNo desc").All()
	if err != nil {
		return nil, err

	}

	return g.Map{
		"items":    all,
		"total":    1,
		"page":     r.Page,
		"pageSize": r.PageSize,
	}, nil
}
func (r *Req) Add() error {

	if _, err := dao.Router.Ctx(r.Ctx).Data(r).Insert(); err != nil {
		return err
	}
	return nil
}

func (r *Req) Edit() error {
	if _, err := dao.Router.Ctx(r.Ctx).Where("id", r.Id).Data(r).Update(); err != nil {
		return err
	}
	return nil
}

func (r *Req) Del() error {
	if _, err := dao.Router.Ctx(r.Ctx).Where("id", r.Id).Delete(); err != nil {
		return err
	}
	return nil
}

func (r *Req) Options() (g.Map, error) {
	panic("implement me")
}

type Routers []entity.Router

func (r *Req) Tree() (g.Map, error) {

	var routers Routers
	if err := g.DB().Model("router").Where("status", 1).Order("parent").Scan(&routers); err != nil {
		return nil, err
	}
	//res := map[int]*Router{}
	//result = BuildRouter(routers)
	generateTree := tree.GenerateTree(routers.ConvertToINodeArray(), nil)

	return g.Map{"router": generateTree}, nil
}

func (p Routers) ConvertToINodeArray() (nodes []tree.INode) {
	for _, v := range p {
		nodes = append(nodes, v)
	}
	return
}
