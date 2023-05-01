package user

import (
	"Gf-Vben/internal/dao"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/jinmao88/gf-utility/curd"
)

type Req struct {
	Ctx context.Context
	curd.Pagination
	Query
}
type Query struct {
	Id       int `p:"id"`
	Username int `p:"username"`
	Uid      int `p:"uid"`
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

	all, err := m.FieldsEx(dao.User.Columns().Password).All()
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
	panic("implement me")
}

func (r *Req) Edit() error {
	panic("implement me")
}

func (r *Req) Del() error {
	panic("implement me")
}
func (r *Req) Tree() (g.Map, error) {
	panic("implement me")
}

func (r *Req) Options() ([]curd.Option, error) {
	panic("implement me")
}
