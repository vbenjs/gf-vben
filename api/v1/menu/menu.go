package menu

import "github.com/gogf/gf/v2/frame/g"

type MenuReq struct {
	g.Meta `path:"/getAll" method:"get" summary:"获取所有路由菜单"`
	Role   string `p:"role"`
}
