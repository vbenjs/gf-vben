package app

import (
	casbin2 "Gf-Vben/app/service/casbin"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	adapter "github.com/jinmao88/gf-casbin"
	"github.com/yitter/idgenerator-go/idgen"
)

// 应用初始化。
func init() {
	//初始化路由
	initRouter()
	//初始化雪花ID
	initIdGenerator()
	//初始化Casbin
	initCasbin()
}

func initIdGenerator() {
	// 创建 IdGeneratorOptions 对象，请在构造函数中输入 WorkerId：
	var options = idgen.NewIdGeneratorOptions(1)
	// options.WorkerIdBitLength = 10 // WorkerIdBitLength 默认值6，支持的 WorkerId 最大值为2^6-1，若 WorkerId 超过64，可设置更大的 WorkerIdBitLength
	// ...... 其它参数设置参考 IdGeneratorOptions 定义，一般来说，只要再设置 WorkerIdBitLength （决定 WorkerId 的最大值）。

	// 保存参数（必须的操作，否则以上设置都不能生效）：
	idgen.SetIdGenerator(options)
}

func initCasbin() {
	a, err := adapter.New(g.DB())
	if err != nil {
		glog.Error(gctx.New(), err)
		return
	}
	modelFromString, err := model.NewModelFromString(`
		[request_definition]
		r = sub, dom, obj, act
		
		[policy_definition]
		p = sub, dom, obj, act
		
		[role_definition]
		g = _, _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = g(r.sub, p.sub, r.dom) && r.dom == p.dom && r.obj == p.obj && (r.act == p.act||p.act == "*")||p.sub == "admin"
	`)
	if err != nil {
		glog.Error(gctx.New(), err)
		return
	}
	casbin2.CE, err = casbin.NewEnforcer(modelFromString, a)
	if err != nil {
		glog.Error(gctx.New(), err)
		return
	}
	domain := casbin2.CE.GetPermissionsForUserInDomain("user", "curd")
	fmt.Println(domain)
	glog.Printf(gctx.New(), "Cabin初始化成功")
}
