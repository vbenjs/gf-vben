package curd

import (
	"Gf-Vben/internal/dao"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
)

func CheckUserId(user int) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		return m.Where(dao.User.Columns().Id, user)
	}
}

func CheckUser(user int) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		return m.Where("user", user)
	}
}

func Order(order string, col ...string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		if order == "" {
			return m
		}
		c := "create_at"
		if len(col) > 0 {
			c = col[0]
		}
		return m.Order(c, order)
	}
}

func TimeRange(s, e any) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		return m.WhereBetween("create_at", s, e)
	}
}

func TimeToday() func(m *gdb.Model) *gdb.Model {
	return TimeRange(gtime.Now().StartOfDay(), gtime.Now())
}
