package boot

import (
	casbin2 "Gf-Vben/app/service/casbin"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	gdbadapter "github.com/vance-liu/gdb-adapter"
	"github.com/yitter/idgenerator-go/idgen"
)

// 用于应用初始化。
func init() {
	//s := g.Server()
	initIdGenerator()
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
	r = sub, dom, obj, act
	
	[policy_definition]
	p = sub, dom, obj, act
	
	[role_definition]
	g = _, _, _
	
	[policy_effect]
	e = some(where (p.eft == allow))
	
	[matchers]
	m = g(r.sub, p.sub, r.dom) && r.dom == p.dom && r.obj == p.obj && r.act == p.act
`)
	if err != nil {
		return
	}
	casbin2.CE, err = casbin.NewEnforcer(modelFromString, a)
	if err != nil {
		glog.Println(err)
		return
	}
	glog.Println("Cabin初始化成功")
}
