package middleware

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

// 初始化中间件
func init() {
	//初始化Jwt
	initJwt()
}

func ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}
	res := r.GetHandlerResponse()
	err := r.GetError()
	if err != nil {
		r.Response.WriteJson(map[string]interface{}{
			"code":    gerror.Code(err).Code(),
			"message": err.Error(),
		})

		return
	}
	m := gconv.Map(res)
	m["data"] = m["result"]
	delete(m, "result")
	r.Response.WriteJson(m)

}
