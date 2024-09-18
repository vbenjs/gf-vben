package middleware

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

// 初始化中间件
func init() {
}

type Resp struct {
	Code int         `json:"code" p:"code"`
	Data interface{} `json:"result" p:"result"`
	Msg  string      `json:"message" p:"message"`
}

func ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()
	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}
	res := r.GetHandlerResponse()
	err := r.GetError()
	resp := new(Resp)
	if err != nil {
		resp.Code = gerror.Code(err).Code()
		resp.Msg = err.Error()
		r.Response.WriteJson(resp)
		return
	}
	gconv.Scan(res, &resp)
	r.Response.WriteJson(resp)

}
