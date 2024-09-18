package middleware

import (
	"Gf-Vben/internal/dao"
	"Gf-Vben/internal/model"
	"Gf-Vben/internal/model/entity"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/jinmao88/gf-utility/response"
)

var GfTokenInstance *gtoken.GfToken

func Login(r *ghttp.Request) (string, interface{}) {
	req := new(model.LoginReq)
	if err := r.Parse(req); err != nil {
		response.JsonExit(r, -1, "参数错误", err)

	}
	var u entity.User
	err := dao.User.Ctx(r.GetCtx()).Where("username", req.Username).Scan(&u)
	if err != nil {
		response.JsonExit(r, -1, "用户或密码不正确")

	}
	if u.Status == 0 {
		response.JsonExit(r, -2, "用户已禁用")
	}
	pw, err := gmd5.Encrypt(req.Password)
	if err != nil {
		response.JsonExit(r, -3, err.Error())

	}
	if pw != u.Password {
		response.JsonExit(r, -1, "用户或密码不正确")
	}
	return u.Username, g.Map{
		"username": u.Username,
		"uid":      u.Id,
		"role":     u.Role,
		"roles":    u.Roles,
	}
}

func LoginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	if respData.Success() {
		response.JsonExit(r, 0, "登录成功", g.Map{
			"accessToken": respData.GetString("token"),
		})
	}
	response.JsonExit(r, -1, respData.Msg, respData.Data)
}
func AuthAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	if respData.Success() {
		data := GfTokenInstance.GetTokenData(r).GetString("data")
		r.SetParamMap(gconv.Map(data))
		r.Middleware.Next()
		return
	}
	response.JsonExit(r, respData.Code, respData.Msg, respData.Data)

}
