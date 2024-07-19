// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-05-10 09:06:44
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HanAdminUserRoleDao is the data access object for table han_admin_user_role.
type HanAdminUserRoleDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns HanAdminUserRoleColumns // columns contains all the column names of Table for convenient usage.
}

// HanAdminUserRoleColumns defines and stores column names for table han_admin_user_role.
type HanAdminUserRoleColumns struct {
	Id         string // 用户关系标识
	Uid        string // 用户id
	Rid        string // 角色id
	CreateTime string // 创建时间
}

// hanAdminUserRoleColumns holds the columns for table han_admin_user_role.
var hanAdminUserRoleColumns = HanAdminUserRoleColumns{
	Id:         "id",
	Uid:        "uid",
	Rid:        "rid",
	CreateTime: "create_time",
}

// NewHanAdminUserRoleDao creates and returns a new DAO object for table data access.
func NewHanAdminUserRoleDao() *HanAdminUserRoleDao {
	return &HanAdminUserRoleDao{
		group:   "default",
		table:   "han_admin_user_role",
		columns: hanAdminUserRoleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HanAdminUserRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HanAdminUserRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HanAdminUserRoleDao) Columns() HanAdminUserRoleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HanAdminUserRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HanAdminUserRoleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HanAdminUserRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
