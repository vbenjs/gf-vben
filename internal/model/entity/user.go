// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id       int         `json:"id"       orm:"id"        ` //
	Username string      `json:"username" orm:"username"  ` //
	Password string      `json:"password" orm:"password"  ` //
	Status   int         `json:"status"   orm:"status"    ` //
	CreateAt *gtime.Time `json:"createAt" orm:"create_at" ` //
	UpdateAt *gtime.Time `json:"updateAt" orm:"update_at" ` //
	DeleteAt *gtime.Time `json:"deleteAt" orm:"delete_at" ` //
	Role     int         `json:"role"     orm:"role"      ` //
}
