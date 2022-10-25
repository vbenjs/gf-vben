package entity

import "github.com/gogf/gf/v2/util/gconv"

// Options 接口实现

func (p Permission) GetOptionLabel() string {
	return p.Name
}

func (p Permission) GetOptionValue() string {
	return gconv.String(p.Id)
}

func (p Permission) GetOptionKey() string {
	return gconv.String(p.Id)
}
