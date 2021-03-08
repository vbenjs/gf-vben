package curd

import "github.com/gogf/gf/frame/g"

type Curd interface {
	List() (g.Map, error)
	Add() error
	Edit() error
	Del() error
}
