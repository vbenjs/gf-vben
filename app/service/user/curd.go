package user

import "github.com/gogf/gf/frame/g"

type Req struct {
	Page     int `p:"page"`
	PageSize int `p:"page_size"`
	Query
}
type Query struct {
	Id       int `p:"id"`
	Username int `p:"username"`
}

func (r Req) List() (g.Map, error) {
	panic("implement me")
}

func (r Req) Add() error {
	panic("implement me")
}

func (r Req) Edit() error {
	panic("implement me")
}

func (r Req) Del() error {
	panic("implement me")
}
