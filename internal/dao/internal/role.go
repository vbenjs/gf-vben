// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoleDao is the data access object for table role.
type RoleDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns RoleColumns // columns contains all the column names of Table for convenient usage.
}

// RoleColumns defines and stores column names for table role.
type RoleColumns struct {
	Id         string //
	Name       string //
	Status     string //
	CreateAt   string //
	UpdateAt   string //
	DeleteAt   string //
	Permission string // 权限集合
}

// roleColumns holds the columns for table role.
var roleColumns = RoleColumns{
	Id:         "id",
	Name:       "name",
	Status:     "status",
	CreateAt:   "create_at",
	UpdateAt:   "update_at",
	DeleteAt:   "delete_at",
	Permission: "permission",
}

// NewRoleDao creates and returns a new DAO object for table data access.
func NewRoleDao() *RoleDao {
	return &RoleDao{
		group:   "default",
		table:   "role",
		columns: roleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RoleDao) Columns() RoleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RoleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
