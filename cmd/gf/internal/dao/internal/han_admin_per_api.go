// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-05-10 09:06:44
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HanAdminPerApiDao is the data access object for table han_admin_per_api.
type HanAdminPerApiDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns HanAdminPerApiColumns // columns contains all the column names of Table for convenient usage.
}

// HanAdminPerApiColumns defines and stores column names for table han_admin_per_api.
type HanAdminPerApiColumns struct {
	Id         string //
	PerId      string //
	Path       string //
	CreateTime string // 创建时间
}

// hanAdminPerApiColumns holds the columns for table han_admin_per_api.
var hanAdminPerApiColumns = HanAdminPerApiColumns{
	Id:         "id",
	PerId:      "per_id",
	Path:       "path",
	CreateTime: "create_time",
}

// NewHanAdminPerApiDao creates and returns a new DAO object for table data access.
func NewHanAdminPerApiDao() *HanAdminPerApiDao {
	return &HanAdminPerApiDao{
		group:   "default",
		table:   "han_admin_per_api",
		columns: hanAdminPerApiColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HanAdminPerApiDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HanAdminPerApiDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HanAdminPerApiDao) Columns() HanAdminPerApiColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HanAdminPerApiDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HanAdminPerApiDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HanAdminPerApiDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
