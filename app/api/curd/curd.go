package curd

import (
	"Gf-Vben/app/service/curd"
	"Gf-Vben/app/service/router"
	"Gf-Vben/app/service/user"
	"Gf-Vben/app/util"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

//import (
//	"Gf-Vben/app/service/curd"
//	"Gf-Vben/app/service/response"
//	"Gf-Vben/app/service/role"
//	"Gf-Vben/app/service/router"
//	"Gf-Vben/app/service/user"
//	"github.com/gogf/gf/v2/net/ghttp"
//)
//

type Api struct {
}

type CurdReq struct {
	g.Meta    `path:"/curd" method:"post" summary:"执行登录请求" tags:"登录"`
	Interface string `p:"i" v:"required"`
	Action    string `p:"a" v:"required"`
}

func (Api) Curd(ctx context.Context, req *CurdReq) (res *util.JsonRes, err error) {
	var cu curd.Curd
	r := g.RequestFromCtx(ctx)
	res = new(util.JsonRes)
	switch req.Interface {
	case "user":
		cu = new(user.Req)
	case "router":
		cu = new(router.Req)
	//case "role":
	//	//req := new(role.Req)
	//	//cu = req
	default:
		return nil, gerror.NewCode(util.Code(1), "接口参数错误")
	}
	if err := r.Parse(cu); err != nil {
		return nil, gerror.NewCode(util.Code(2), err.Error())
	}
	switch req.Action {
	case "list":
		result, err := cu.List()
		if err != nil {
			return nil, gerror.NewCode(util.Code(3), err.Error())
		}
		res.Data = result
	case "tree":
		result, err := cu.Tree()
		if err != nil {
			return nil, gerror.NewCode(util.Code(3), err.Error())
		}
		res.Data = result
	case "options":
		result, err := cu.Options()
		if err != nil {
			return nil, gerror.NewCode(util.Code(3), err.Error())
		}
		res.Data = result
	case "add":
		if err := cu.Add(); err != nil {
			return nil, gerror.NewCode(util.Code(3), err.Error())
		}
		res.Message = "新增成功"
	case "edit":
		if err := cu.Edit(); err != nil {
			return nil, gerror.NewCode(util.Code(3), err.Error())
		}
		res.Message = "修改成功"
	case "del":
		if err := cu.Del(); err != nil {
			return nil, gerror.NewCode(util.Code(3), err.Error())
		}
		res.Message = "修改成功"
	default:
		return nil, gerror.NewCode(util.Code(3), "接口参数错误")
	}
	return

}
