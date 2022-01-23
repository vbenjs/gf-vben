package user

import (
	"Gf-Vben/app/service/internal/dao"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
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
	Uuid     int `p:"uuid"`
}

func (r *Req) List() (g.Map, error) {

	u, err := dao.User.Ctx(r.Ctx).All()
	if err != nil {
		return nil, err
	}

	if u == nil {
		return nil, gerror.New("用户不存在")
	}
	return g.Map{
		"items":    u,
		"total":    1,
		"page":     r.Page,
		"pageSize": r.PageSize,
	}, nil
	//return g.Map{"list":[]g.Map{g.Map{}}}
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

func (r *Req) Options() (g.Map, error) {
	panic("implement me")
}
