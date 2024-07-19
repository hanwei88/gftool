// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-05-10 09:06:44
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HanAdminGroupRoleDao is the data access object for table han_admin_group_role.
type HanAdminGroupRoleDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns HanAdminGroupRoleColumns // columns contains all the column names of Table for convenient usage.
}

// HanAdminGroupRoleColumns defines and stores column names for table han_admin_group_role.
type HanAdminGroupRoleColumns struct {
	Id         string // 分组标识
	Gid        string // 分组id
	Rid        string // 角色id
	CreateTime string // 创建时间
}

// hanAdminGroupRoleColumns holds the columns for table han_admin_group_role.
var hanAdminGroupRoleColumns = HanAdminGroupRoleColumns{
	Id:         "id",
	Gid:        "gid",
	Rid:        "rid",
	CreateTime: "create_time",
}

// NewHanAdminGroupRoleDao creates and returns a new DAO object for table data access.
func NewHanAdminGroupRoleDao() *HanAdminGroupRoleDao {
	return &HanAdminGroupRoleDao{
		group:   "default",
		table:   "han_admin_group_role",
		columns: hanAdminGroupRoleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HanAdminGroupRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HanAdminGroupRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HanAdminGroupRoleDao) Columns() HanAdminGroupRoleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HanAdminGroupRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HanAdminGroupRoleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HanAdminGroupRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
