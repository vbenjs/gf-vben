package curd

import (
	"Gf-Vben/app/service/curd"
	"Gf-Vben/app/service/permission"
	"Gf-Vben/app/service/role"
	"Gf-Vben/app/service/router"
	"Gf-Vben/app/service/user"
	"Gf-Vben/app/util"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type Api struct {
}

type CurdReq struct {
	g.Meta    `path:"/curd" method:"post" summary:"Curd请求" tags:"Curd"`
	Interface string `p:"i" v:"required"`
	Action    string `p:"a" v:"required"`
}

func (Api) Curd(ctx context.Context, req *CurdReq) (res *util.JsonRes, err error) {
	var cu curd.Curd
	res = new(util.JsonRes)
	switch req.Interface {
	case "user":
		cu = new(user.Req)
	case "router":
		cu = new(router.Req)
	case "role":
		cu = new(role.Req)
	case "permission":
		cu = new(permission.Req)
	//	//cu = req
	default:
		return nil, gerror.NewCode(util.Code(1), "接口参数错误")
	}
	if err = g.RequestFromCtx(ctx).Parse(cu); err != nil {
		return nil, gerror.NewCode(util.Code(2), err.Error())
	}
	switch req.Action {
	case "list":
		res.Data, err = cu.List()
	case "tree":
		res.Data, err = cu.Tree()
	case "options":
		res.Data, err = cu.Options()
	case "add":
		err = cu.Add()
		res.Message = "新增成功"
	case "edit":
		err = cu.Edit()
		res.Message = "修改成功"
	case "del":
		err = cu.Del()
		res.Message = "删除成功"
	default:
		err = gerror.NewCode(util.Code(3), "接口参数错误")
	}
	return

}
