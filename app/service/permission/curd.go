package permission

import (
	"Gf-Vben/app/model/entity"
	"Gf-Vben/app/service/internal/dao"
	"context"
	tree "github.com/azhengyongqin/golang-tree-menu"
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
	Tag    string `p:"tag"`
	Parent int    `p:"parent"`
	Desc   string `p:"desc"`
}

type Role struct {
	Name string `json:"name"`
}
type Permissions []entity.Permission

func (r *Req) List() (g.Map, error) {
	var res Permissions

	if err := dao.Permission.Ctx(r.Ctx).Scan(&res); err != nil {
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
	panic("implement me")
}

func (r *Req) Edit() error {
	panic("implement me")
}

func (r *Req) Del() error {
	panic("implement me")
}

func (r *Req) Tree() (g.Map, error) {
	var res Permissions
	if err := dao.Permission.Ctx(r.Ctx).Scan(&res); err != nil {
		return nil, err
	}
	generateTree := tree.GenerateTree(res.ConvertToINodeArray(), nil)
	return g.Map{"tree": generateTree}, nil
}

func (r *Req) Options() (g.Map, error) {
	panic("implement me")
}

func (p Permissions) ConvertToINodeArray() (nodes []tree.INode) {
	for _, v := range p {
		nodes = append(nodes, v)
	}
	return
}
