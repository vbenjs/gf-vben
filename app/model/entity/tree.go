package entity

import "github.com/gogf/gf/v2/util/gconv"

func (p Permission) GetTitle() string {
	return p.Tag
}

func (p Permission) GetId() int {
	return gconv.Int(p.Id)
}

func (p Permission) GetFatherId() int {
	return gconv.Int(p.Parent)
}

func (p Permission) GetData() interface{} {
	return p
}

func (p Permission) IsRoot() bool {
	return p.Parent == 0
}

func (p Router) GetTitle() string {
	return p.Title
}

func (p Router) GetId() int {
	return gconv.Int(p.Id)
}

func (p Router) GetFatherId() int {
	return gconv.Int(p.Parent)
}

func (p Router) GetData() interface{} {
	return p
}

func (p Router) IsRoot() bool {
	return p.Parent == 0
}
