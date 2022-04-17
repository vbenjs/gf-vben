package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

var CE *casbin.Enforcer

type Req struct {
	Uid       string `p:"uid"`
	Domain    string
	Interface string `p:"i"`
	Action    string `p:"a"`
}

func (r *Req) Check() error {
	g.Dump(r)
	t, err := CE.Enforce(r.Uid, r.Domain, r.Interface, r.Action)
	if err != nil {
		return err
	}
	if !t {
		return gerror.New("无此权限")
	}

	return nil

}
