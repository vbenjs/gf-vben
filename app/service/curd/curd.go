package curd

import (
	"Gf-Vben/app/util/options"
	"github.com/gogf/gf/v2/frame/g"
)

type Curd interface {
	// List 列表
	List() (g.Map, error)
	// Add 新增
	Add() error
	// Edit 编辑
	Edit() error
	// Del 删
	Del() error
	// Tree 返回树结构
	Tree() (g.Map, error)
	// Options 返回options
	Options() ([]options.Option, error)
}
