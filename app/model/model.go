// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package model

import (
	"github.com/gogf/gf/os/gtime"
)

// CasbinRule is the golang structure for table casbin_rule.
type CasbinRule struct {
	Ptype string `orm:"ptype" json:"ptype"` //
	V0    string `orm:"v0"    json:"v0"`    //
	V1    string `orm:"v1"    json:"v1"`    //
	V2    string `orm:"v2"    json:"v2"`    //
	V3    string `orm:"v3"    json:"v3"`    //
	V4    string `orm:"v4"    json:"v4"`    //
	V5    string `orm:"v5"    json:"v5"`    //
}

// Router is the golang structure for table router.
type Router struct {
	Id        int         `orm:"id,primary" json:"id"`        //
	Path      string      `orm:"path"       json:"path"`      //
	Name      string      `orm:"name"       json:"name"`      //
	Redirect  string      `orm:"redirect"   json:"redirect"`  //
	Title     string      `orm:"title"      json:"title"`     //
	Icon      string      `orm:"icon"       json:"icon"`      //
	Component string      `orm:"component"  json:"component"` //
	Parent    int         `orm:"parent"     json:"parent"`    //
	OrderNo   int         `orm:"orderNo"    json:"orderNo"`   //
	Status    int         `orm:"status"     json:"status"`    //
	CreateAt  *gtime.Time `orm:"create_at"  json:"createAt"`  //
	UpdateAt  *gtime.Time `orm:"update_at"  json:"updateAt"`  //
	DeleteAt  *gtime.Time `orm:"delete_at"  json:"deleteAt"`  //
}

// User is the golang structure for table user.
type User struct {
	Id       int         `orm:"id,primary"       json:"id"`       // primary id
	Username string      `orm:"username,primary" json:"username"` // username
	Password string      `orm:"password"         json:"password"` // password
	NickName string      `orm:"nick_name"        json:"nickName"` // nickName
	Status   int         `orm:"status"           json:"status"`   // 1:enable 2:disable
	CreateAt *gtime.Time `orm:"create_at"        json:"createAt"` //
	UpdateAt *gtime.Time `orm:"update_at"        json:"updateAt"` //
	DeleteAt *gtime.Time `orm:"delete_at"        json:"deleteAt"` //
}