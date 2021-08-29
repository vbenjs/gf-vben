package router

import (
	"Gf-Vben/app/dao"
	"github.com/gogf/gf/frame/g"
)

type Req struct {
	Page     int `p:"page"`
	PageSize int `p:"page_size"`
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
		"items":    getRouterList(all),
		"total":    1,
		"page":     r.Page,
		"pageSize": r.PageSize,
	}, nil
}
func (r Req) Add() error {
	if _, err := dao.Router.Data(r).Save(); err != nil {
		return err
	}
	return nil
}

func (r Req) Edit() error {
	if _, err := dao.Router.Where("id", r.Id).Data(r).Save(); err != nil {
		return err
	}
	return nil
}

func (r Req) Del() error {
	if _, err := dao.Router.Where("id", r.Id).Delete(); err != nil {
		return err
	}
	return nil
}
