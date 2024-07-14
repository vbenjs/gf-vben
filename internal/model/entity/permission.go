// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Permission is the golang structure for table permission.
type Permission struct {
	Id       int         `json:"id"       orm:"id"        ` //
	Name     string      `json:"name"     orm:"name"      ` //
	Parent   int         `json:"parent"   orm:"parent"    ` //
	Type     int         `json:"type"     orm:"type"      ` // 权限类型 1、权限域 2、权限组 3、权限操作
	CreateAt *gtime.Time `json:"createAt" orm:"create_at" ` //
	UpdateAt *gtime.Time `json:"updateAt" orm:"update_at" ` //
	DeleteAt *gtime.Time `json:"deleteAt" orm:"delete_at" ` //
}
