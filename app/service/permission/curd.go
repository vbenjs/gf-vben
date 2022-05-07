package permission

import (
	"Gf-Vben/app/model/entity"
	"Gf-Vben/app/service/internal/dao"
	"Gf-Vben/app/util/options"
	"Gf-Vben/app/util/tree"
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
type Permissions []entity.Permission

func (r *Req) List() (g.Map, error) {
	var res Permissions
	m := dao.Permission.Ctx(r.Ctx)
	if r.Type != "" {
		m = m.Where(dao.Permission.Columns().Type, r.Type)
	}

	if err := m.OrderAsc(dao.Permission.Columns().Id).Scan(&res); err != nil {
		return nil, err
	}

	return g.Map{
		"items":    res,
		"total":    len(res),
		"page":     r.Page,
		"pageSize": r.PageSize,
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

	t := tree.GenerateTree(tree.ConvertToINodeArray(res), nil)

	return g.Map{"tree": t}, nil
}

func (r *Req) Options() ([]options.Option, error) {
	var res []entity.Permission
	m := dao.Permission.Ctx(r.Ctx)
	if r.Type != "" {
		m = m.Where(dao.Permission.Columns().Type, r.Type)
	}

	if err := m.OrderAsc(dao.Permission.Columns().Id).Scan(&res); err != nil {
		return nil, err
	}

	return options.BuildOptions(res), nil
}
