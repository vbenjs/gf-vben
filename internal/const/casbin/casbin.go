package casbin

//var Enforcer *casbin.Enforcer

type Req struct {
	Uid       string `p:"uid"`
	Domain    string
	Interface string `p:"i"`
	Action    string `p:"a"`
}

//func (r *Req) Check() error {
//	g.Dump(r)
//	t, err := Enforcer.Enforce(r.Uid, r.Domain, r.Interface, r.Action)
//	if err != nil {
//		return err
//	}
//	if !t {
//		return gerror.New("无此权限")
//	}
//
//	return nil
//
//}
