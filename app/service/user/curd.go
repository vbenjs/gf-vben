package user

import (
	"Gf-Vben/app/model/entity"
	"Gf-Vben/app/service/curd"
	"Gf-Vben/app/service/internal/dao"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

type Req struct {
	Ctx      context.Context
	Page     int `p:"page"`
	PageSize int `p:"page_size"`
	Query
}
type Query struct {
	Id       int `p:"id"`
	Username int `p:"username"`
	Uid      int `p:"uid"`
}

func (r *Req) List() (*curd.List, error) {
	res := make([]entity.User, 0)
	if err := dao.User.Ctx(r.Ctx).Scan(&res); err != nil {
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
