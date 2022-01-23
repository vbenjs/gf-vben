package role

import (
	"Gf-Vben/app/service/casbin"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
)

type Req struct {
	Page     int `p:"page"`
	PageSize int `p:"page_size"`
	Query
}

type Query struct {
	Id       int `p:"id"`
	Username int `p:"username"`
	Uuid     int `p:"uuid"`
}

type Role struct {
	Name string `json:"name"`
}

func (r *Req) List() (g.Map, error) {
	roles := casbin.CE.GetAllRoles()
	var res []Role
	for _, role := range roles {
		res = append(res, Role{Name: role})
		domain, _ := casbin.CE.GetDomainsForUser(role)
		inDomain := casbin.CE.GetPermissionsForUserInDomain(role, "curd")
		casbin.CE.GetRoleManager().PrintRoles()
		fmt.Println(domain)
		fmt.Println(inDomain)
	}
	return g.Map{
		"items":    res,
		"total":    1,
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
	panic("implement me")
}

func (r *Req) Options() (g.Map, error) {
	panic("implement me")
}
