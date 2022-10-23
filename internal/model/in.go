package model

type CurdReq struct {
	Interface string `p:"i" v:"required"`
	Action    string `p:"a" v:"required"`
}
type LoginReq struct {
	Username string `json:"username" dc:"账号" v:"required"`
	Password string `json:"password" dc:"密码" v:"required"`
}

type RegisterReq struct {
	Username string `p:"username" v:"required"`
	Pw       string `p:"password" v:"required"`
	Pw2      string `p:"password2" v:"required"`
}
