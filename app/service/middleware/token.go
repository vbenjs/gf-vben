package middleware

//
//import (
//	"Gf-Vben/app/dao"
//	"Gf-Vben/app/model"
//	"Gf-Vben/app/service/response"
//	"github.com/goflyfox/gtoken/gtoken"
//	"/v2/crypto/gmd5"
//	"github.com/gogf/gf/v2/frame/g"
//	"github.com/gogf/gf/v2/net/ghttp"
//	"github.com/gogf/gf/v2/util/gconv"
//)
//
//var Gtoken = &gtoken.GfToken{
//	LoginPath:        "/login",
//	LoginBeforeFunc:  loginFunc,
//	LoginAfterFunc:   loginAfterFunc,
//	LogoutPath:       "/api/user/logout",
//	AuthAfterFunc:    authAfter,
//	AuthPaths:        g.SliceStr{"/api"},
//	GlobalMiddleware: true,
//}
//
//func loginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
//	response.JsonExit(r, 0, "", respData.Data)
//}
//
//func authAfter(r *ghttp.Request, respData gtoken.Resp) {
//	data := gconv.Map(respData.GetString("data"))
//	if data == nil {
//		response.JsonExit(r, 401, "未登录")
//	}
//	r.SetParam("uid", data["uid"])
//	r.Middleware.Next()
//}
//
//func loginFunc(r *ghttp.Request) (string, interface{}) {
//	req := new(LoginReq)
//	if err := r.Parse(req); err != nil {
//		response.JsonExit(r, 1, err.Error())
//	}
//	var u model.User
//	err := dao.User.Where("username", req.Username).Scan(&u)
//	if err != nil {
//		response.JsonExit(r, 2, err.Error())
//	}
//	if u.Status == 0 {
//		response.JsonExit(r, 3, "用户已禁用")
//
//	}
//	pw, err := gmd5.Encrypt(req.Password)
//	if err != nil {
//		response.JsonExit(r, 3, err.Error())
//
//	}
//	if pw != u.Password {
//		response.JsonExit(r, 3, "用户密码错误")
//	}
//	return u.Username, g.Map{"uid": u.Id, "roles": u.Username}
//}
