package router

import (
	"Gf-Vben/internal/model/entity"
	"github.com/gogf/gf/v2/util/gconv"
)

type Entity struct {
	entity.Router
}

func (e Entity) GetTitle() string {
	return e.Title
}

func (e Entity) GetId() int {
	return gconv.Int(e.Id)
}

func (e Entity) GetFatherId() int {
	return gconv.Int(e.Parent)
}

func (e Entity) GetData() interface{} {
	return e
}

func (e Entity) IsRoot() bool {
	return e.Parent == 0
}

func (e Entity) GetTreeValue() interface{} {
	return e.Id == 0
}
