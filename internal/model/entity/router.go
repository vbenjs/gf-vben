// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Router is the golang structure for table router.
type Router struct {
	Id        int         `json:"id"        orm:"id"        ` //
	Path      string      `json:"path"      orm:"path"      ` //
	Name      string      `json:"name"      orm:"name"      ` //
	Redirect  string      `json:"redirect"  orm:"redirect"  ` //
	Title     string      `json:"title"     orm:"title"     ` //
	Icon      string      `json:"icon"      orm:"icon"      ` //
	Component string      `json:"component" orm:"component" ` //
	Parent    int         `json:"parent"    orm:"parent"    ` //
	OrderNo   int         `json:"orderNo"   orm:"order_no"  ` //
	Status    int         `json:"status"    orm:"status"    ` //
	CreateAt  *gtime.Time `json:"createAt"  orm:"create_at" ` //
	UpdateAt  *gtime.Time `json:"updateAt"  orm:"update_at" ` //
	DeleteAt  *gtime.Time `json:"deleteAt"  orm:"delete_at" ` //
}
