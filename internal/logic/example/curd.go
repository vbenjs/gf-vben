package example

import (
	"Gf-Vben/internal/const/curd"
	"Gf-Vben/internal/dao"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

type Req struct {
	curd.Pagination
	Ctx context.Context `p:"ctx"`
	Query
}

type Query struct {
	Id int `p:"id"`
}

func (r *Req) SetCtx(ctx context.Context) {
	r.Ctx = ctx
}

func (r *Req) List() (*curd.List, error) {
	m := dao.User.Ctx(r.Ctx)
	count, err := m.Count()
	if err != nil {
		return nil, err
	}
	all, err := m.Page(r.Page, r.PageSize).All()
	if err != nil {
		return nil, err
	}
	return &curd.List{
		Items:    all,
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
	//var res []entity.Permission
	//if err := dao.Permission.Ctx(r.Ctx).Scan(&res); err != nil {
	//	return nil, err
	//}
	//generateTree := curd.GenerateTree(res, nil)
	return g.Map{"tree": ""}, nil
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
