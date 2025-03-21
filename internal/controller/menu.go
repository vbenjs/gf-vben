package controller

import (
	"Gf-Vben/api/v1/menu"
	"Gf-Vben/internal/middleware"
	"Gf-Vben/internal/service"
	"context"
	"github.com/jinmao88/gf-utility/response"
)

var (
	Menu = cMenu{}
)

type cMenu struct {
}

func (cMenu) All(ctx context.Context, req *menu.MenuReq) (res response.JsonRes, err error) {
	r := new(middleware.Resp)
	r.Data, err = service.User().Menu(ctx)
	return r, nil
}
