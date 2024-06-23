package router

import (
	"Gf-Vben/internal/dao"
	"Gf-Vben/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/jinmao88/gf-utility/curd"
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

func (r *Req) SetCtx(ctx context.Context) {
	r.Ctx = ctx
}

func (r *Req) List() (*curd.List, error) {
	res := make([]entity.Router, 0)
	if err := dao.Router.Ctx(r.Ctx).Order("parent,orderNo desc").Scan(&res); err != nil {
		return nil, err
	}

	return &curd.List{
		Items:    res,
		Total:    len(res),
		Page:     r.Page,
		PageSize: r.PageSize,
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

func (r *Req) Options() ([]curd.Option, error) {
	panic("implement me")
}

func (r *Req) Tree() (g.Map, error) {

	var res []Entity
	if err := g.DB().Model("router").Where("status", 1).Order("parent").Scan(&res); err != nil {
		return nil, err
	}
	//res := map[int]*Router{}
	//result = BuildRouter(routers)
	generateTree := curd.GenerateTree(res, nil)

	return g.Map{"router": generateTree}, nil
}
