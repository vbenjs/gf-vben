package example

import (
	"Gf-Vben/app/model/entity"
	"Gf-Vben/app/service/curd"
	"Gf-Vben/app/service/internal/dao"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

type Req struct {
	Page     int             `p:"page"`
	PageSize int             `p:"page_size"`
	Ctx      context.Context `p:"ctx"`
	Query
}

type Query struct {
	Id     int    `p:"id"`
	User   int    `p:"user"`
	Name   string `p:"name"`
	NameEn string `p:"name_en"`
	Desc   string `p:"desc"`
	DescEn string `p:"desc_en"`
	Icon   string `p:"icon"`
	Href   string `p:"href"`
	Sort   int    `p:"sort"`
	Type   int    `p:"type"`
	Status string `p:"status"`
}

func (r *Req) List() (*curd.List, error) {
	res := make([]entity.User, 0)
	m := dao.User.Ctx(r.Ctx)
	count, err := m.Count()
	if err != nil {
		return nil, err
	}
	if err := m.Page(r.Page, r.PageSize).Scan(res); err != nil {
		return nil, err
	}
	return &curd.List{
		Items:    res,
		Total:    count,
		Page:     r.Page,
		PageSize: r.PageSize,
	}, nil
}

func (r *Req) Add() error {

	return nil
}

func (r *Req) Edit() error {

	return nil

}

func (r *Req) Del() error {
	panic("implement me")
}

func (r *Req) Tree() (g.Map, error) {
	var res []entity.Permission
	if err := dao.Permission.Ctx(r.Ctx).Scan(&res); err != nil {
		return nil, err
	}
	generateTree := curd.GenerateTree(curd.ConvertToINodeArray(res), nil)
	return g.Map{"tree": generateTree}, nil
}

func (r *Req) Options() ([]curd.Option, error) {
	//var res []entity.Permission
	//m := dao.Permission.Ctx(r.Ctx)
	//if r.Type != "" {
	//	m = m.Where(dao.Permission.Columns().Type, r.Type)
	//}
	//
	//if err := m.OrderAsc(dao.Permission.Columns().Id).Scan(&res); err != nil {
	//	return nil, err
	//}
	//
	//return options.BuildOptions(res), nil
	return []curd.Option{}, nil
}
