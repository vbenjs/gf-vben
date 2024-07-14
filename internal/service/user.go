// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"Gf-Vben/internal/model"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/jinmao88/gf-utility/menu"
)

type (
	IUser interface {
		Register(ctx context.Context, in model.RegisterReq) error
		Menu(ctx context.Context) ([]*menu.Menu, error)
		Info(ctx context.Context, uid int) (gdb.Record, error)
		AccessCode(ctx context.Context, role int) ([]string, error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
