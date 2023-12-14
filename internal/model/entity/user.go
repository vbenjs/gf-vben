// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id       int         `json:"id"       ` // primary id
	Username string      `json:"username" ` // username
	Password string      `json:"password" ` // password
	Note     string      `json:"note"     ` //
	NickName string      `json:"nickName" ` // nickName
	Status   int         `json:"status"   ` // 1:enable 2:disable
	CreateAt *gtime.Time `json:"createAt" ` //
	UpdateAt *gtime.Time `json:"updateAt" ` //
	DeleteAt *gtime.Time `json:"deleteAt" ` //
}
