package user

//
//import (
//	"github.com/gogf/gf/frame/g"
//	"init/app/dao"
//)
//
//type Req struct {
//	Page     int `p:"page"`
//	PageSize int `p:"page_size"`
//	Query
//}
//type Query struct {
//	Id     int    `p:"id"`
//	User   int    `p:"user"`
//	Name   string `p:"name"`
//	NameEn string `p:"name_en"`
//	Desc   string `p:"desc"`
//	DescEn string `p:"desc_en"`
//	Icon   string `p:"icon"`
//	Href   string `p:"href"`
//	Sort   int    `p:"sort"`
//	Type   int    `p:"type"`
//	Status string `p:"status"`
//}
//
//func (r Req) List() (g.Map, error) {
//	m := g.DB().Model("nav")
//	count, err := m.Count()
//	if err != nil {
//		return nil, err
//	}
//	all, err := m.Page(r.Page, r.PageSize).All()
//	if err != nil {
//		return nil, err
//	}
//	return g.Map{
//		"items":    all,
//		"total":    count,
//		"page":     r.Page,
//		"pageSize": r.PageSize,
//	}, nil
//}
//
//func (r Req) Add() error {
//	if _, err := dao.Nav.OmitEmpty().Data(r).Save(); err != nil {
//		return err
//	}
//	return nil
//}
//
//func (r Req) Edit() error {
//	if _, err := dao.Nav.OmitEmpty().Where("id", r.Id).Update(r); err != nil {
//		return err
//	}
//	return nil
//
//}
//
//func (r Req) Del() error {
//	panic("implement me")
//}
