package curd

import "github.com/gogf/gf/frame/g"

type Curd interface {
	// List 列表
	List() (g.Map, error)
	// Add 新增
	Add() error
	// Edit 编辑
	Edit() error
	// Del 删除
	Del() error
	// Tree 返回树结构
	Tree() (g.Map, error)
	// Options 返回options
	Options() (g.Map, error)
}
