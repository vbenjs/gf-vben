package boot

import (
	"Gf-Vben/app/service/middleware"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	gdbadapter "github.com/vance-liu/gdb-adapter"
)

// 用于应用初始化。
func init() {
	//s := g.Server()
	initCasbin()
}

func initCasbin() {
	opts := &gdbadapter.Adapter{
		// or reuse an existing connection:
		Db: g.DB(),
	}
	a, err := gdbadapter.NewAdapterFromOptions(opts)
	if err != nil {
		glog.Println(err)
		return
	}

	modelFromString, err := model.NewModelFromString(`
	[request_definition]
	r = sub, obj, act
	
	[policy_definition]
	p = sub, obj, act
	
	[role_definition]
	g = _, _
	
	[policy_effect]
	e = some(where (p.eft == allow))
	
	[matchers]
	m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
`)
	if err != nil {
		return
	}
	middleware.CE, err = casbin.NewEnforcer(modelFromString, a)
	if err != nil {
		glog.Println(err)
		return
	}
	glog.Println("Cabin初始化成功")
}
