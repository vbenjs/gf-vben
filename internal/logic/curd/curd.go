package curd

import (
	"Gf-Vben/internal/const/curd"
	"Gf-Vben/internal/logic/permission"
	"Gf-Vben/internal/logic/role"
	"Gf-Vben/internal/logic/router"
	"Gf-Vben/internal/logic/user"
	"Gf-Vben/internal/model"
	"Gf-Vben/internal/service"
	"Gf-Vben/util"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sCurd struct {
}

func init() {
	service.RegisterCurd(New())
}
func New() *sCurd {
	return &sCurd{}
}

func (s *sCurd) Curd(ctx context.Context, r *model.CurdReq) (res *util.JsonRes, err error) {
	var cu curd.Curd

	res = new(util.JsonRes)
	switch r.Interface {
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
	switch r.Action {
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
