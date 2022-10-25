package permission

import (
	"Gf-Vben/internal/const/curd"
	"Gf-Vben/internal/dao"
	"Gf-Vben/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type Req struct {
	Page     int `p:"page"`
	PageSize int `p:"page_size"`
	Ctx      context.Context
	Query
}

type Query struct {
	Id     int    `p:"id"`
	Uid    int    `p:"uid"`
	Name   string `p:"name"`
	Value  string `p:"value"`
	Parent int    `p:"parent"`
	Type   string `p:"type"`
	Desc   string `p:"desc"`
}

type Role struct {
	Name string `json:"name"`
}

func (r *Req) SetCtx(ctx context.Context) {
	r.Ctx = ctx
}

type Permissions []entity.Permission

func (r *Req) List() (*curd.List, error) {
	var res Permissions
	m := dao.Permission.Ctx(r.Ctx)
	if r.Type != "" {
		m = m.Where(dao.Permission.Columns().Type, r.Type)
	}

	if err := m.OrderAsc(dao.Permission.Columns().Id).Scan(&res); err != nil {
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
	if _, err := dao.Permission.Ctx(r.Ctx).OmitEmptyData().Insert(r); err != nil {
		return err
	}

	return nil
}

func (r *Req) Edit() error {
	if _, err := dao.Permission.Ctx(r.Ctx).OmitEmptyData().Where(dao.Permission.Columns().Id, r.Id).Update(r); err != nil {
		return err
	}
	return nil
}

func (r *Req) Del() error {
	var p entity.Permission
	if err := dao.Permission.Ctx(r.Ctx).Where(dao.Permission.Columns().Id, r.Id).Scan(&p); err != nil {
		return err
	}
	if p.Id == 0 {
		return gerror.New("记录不存在")
	}
	if p.Parent == 0 {
		if _, err := dao.Permission.Ctx(r.Ctx).Where(dao.Permission.Columns().Parent, p.Id).Delete(); err != nil {
			return err
		}

	}
	if _, err := dao.Permission.Ctx(r.Ctx).Where(dao.Permission.Columns().Id, r.Id).Delete(); err != nil {
		return err
	}
	return nil

}

func (r *Req) Tree() (g.Map, error) {
	var res []entity.Permission
	m := dao.Permission.Ctx(r.Ctx)
	if r.Type != "" {
		m = m.WhereNot(dao.Permission.Columns().Type, 3)
	}
	if err := m.Scan(&res); err != nil {
		return nil, err
	}

	t := curd.GenerateTree(res, nil)

	return g.Map{"tree": t}, nil
}

func (r *Req) Options() ([]curd.Option, error) {
	var res []entity.Permission
	m := dao.Permission.Ctx(r.Ctx)
	if r.Type != "" {
		m = m.Where(dao.Permission.Columns().Type, r.Type)
	}

	if err := m.OrderAsc(dao.Permission.Columns().Id).Scan(&res); err != nil {
		return nil, err
	}

	return curd.BuildOptions(res), nil
}
