// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure for table role.
type Role struct {
	Id         int         `json:"id"         orm:"id"         ` //
	Name       string      `json:"name"       orm:"name"       ` //
	Status     bool        `json:"status"     orm:"status"     ` //
	CreateAt   *gtime.Time `json:"createAt"   orm:"create_at"  ` //
	UpdateAt   *gtime.Time `json:"updateAt"   orm:"update_at"  ` //
	DeleteAt   *gtime.Time `json:"deleteAt"   orm:"delete_at"  ` //
	Permission []int       `json:"permission" orm:"permission" ` // 权限集合
}
