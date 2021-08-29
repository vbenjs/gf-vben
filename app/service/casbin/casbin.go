package casbin

import (
	"github.com/casbin/casbin/v2"
)

var CE *casbin.Enforcer

type Req struct {
	Uuid      string `p:"uuid"`
	Interface string `p:"i"`
	Action    string `p:"a"`
}

func (r *Req) Check() error {

	if _, err := CE.Enforce(r.Uuid, "curd", r.Interface, r.Action); err != nil {
		return err
	}
	return nil

}
