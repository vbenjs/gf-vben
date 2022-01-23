package middleware

//func Casbin(r *ghttp.Request) {
//	var req casbin.Req
//	if err := r.Parse(&req); err != nil {
//		response.JsonExit(r, 1, "权限失效")
//	}
//
//	if err := req.Check(); err != nil {
//		response.JsonExit(r, 2, err.Error())
//	}
//	r.Middleware.Next()
//
//}
