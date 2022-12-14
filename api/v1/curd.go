package v1

import (
	"Gf-Vben/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type CurdReq struct {
	g.Meta `path:"/curd" method:"post" summary:"Curd请求" tags:"Curd"`
	model.CurdReq
}
