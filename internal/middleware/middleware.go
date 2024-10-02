package middleware

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/jinmao88/gf-utility/response"
)

// 初始化中间件
func init() {
}

type Resp struct {
	Code    int         `json:"code" p:"code"`
	Data    interface{} `json:"data,omitempty" p:"data"`
	Message string      `json:"message" p:"message"`
}

func (res *Resp) JsonExit(r *ghttp.Request) {
	r.Response.WriteJsonExit(res)
}
func (res *Resp) Format(code int, msg string, data ...interface{}) response.JsonRes {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	return &Resp{
		Code:    code,
		Message: msg,
		Data:    responseData,
	}
}
func (res *Resp) M(msg string) response.JsonRes {
	res.Message = msg
	return res
}

func (res *Resp) D(data interface{}) response.JsonRes {
	res.Data = data
	return res
}

func (res *Resp) C(code int) response.JsonRes {
	res.Code = code
	return res
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
		resp.C(gerror.Code(err).Code()).M(err.Error()).JsonExit(r)
		return
	}
	gconv.Scan(res, &resp)
	r.Response.WriteJson(resp)

}
