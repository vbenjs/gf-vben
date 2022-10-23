package model

type CurdReq struct {
	Interface string `p:"i" v:"required"`
	Action    string `p:"a" v:"required"`
}
