package role

import (
	"Gf-Vben/app/model/entity"
	"Gf-Vben/app/service/internal/dao"
	"Gf-Vben/app/util/options"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

type Req struct {
	Page     int `p:"page"`
	PageSize int `p:"page_size"`
	Ctx      context.Context
	Query
}

type Query struct {
	Id         int    `p:"id"`
	Uid        int    `p:"uid"`
	Name       string `p:"name"`
	Value      string `p:"value"`
	Desc       string `p:"desc"`
	Permission string `p:"permission"`
}

type Role struct {
	Name string `json:"name"`
}

func (r *Req) List() (g.Map, error) {
	var res []entity.Role
	if err := dao.Role.Ctx(r.Ctx).Order("id").Scan(&res); err != nil {
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
	if _, err := dao.Role.Ctx(r.Ctx).Insert(r); err != nil {
		return err
	}
	return nil
}

func (r *Req) Edit() error {

	if _, err := dao.Role.Ctx(r.Ctx).Where("id", r.Id).OmitEmptyData().Update(r); err != nil {
		return err
	}
	return nil
}

func (r *Req) Del() error {
	if _, err := dao.Role.Ctx(r.Ctx).Delete("id", r.Id); err != nil {
		return err
	}
	return nil
}

func (r *Req) Tree() (g.Map, error) {
	panic("implement me")
}

func (r *Req) Options() ([]options.Option, error) {
	panic("implement me")
}
