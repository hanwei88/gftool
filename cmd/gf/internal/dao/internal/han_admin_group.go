// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-05-10 09:06:44
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HanAdminGroupDao is the data access object for table han_admin_group.
type HanAdminGroupDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns HanAdminGroupColumns // columns contains all the column names of Table for convenient usage.
}

// HanAdminGroupColumns defines and stores column names for table han_admin_group.
type HanAdminGroupColumns struct {
	Id         string // 分组标识
	Name       string // 分组名称
	Type       string // 类型 0：分组或者部门 1：职位
	Pid        string // 上下级
	CreateTime string // 创建时间
	DeleteTime string // 删除时间
}

// hanAdminGroupColumns holds the columns for table han_admin_group.
var hanAdminGroupColumns = HanAdminGroupColumns{
	Id:         "id",
	Name:       "name",
	Type:       "type",
	Pid:        "pid",
	CreateTime: "create_time",
	DeleteTime: "delete_time",
}

// NewHanAdminGroupDao creates and returns a new DAO object for table data access.
func NewHanAdminGroupDao() *HanAdminGroupDao {
	return &HanAdminGroupDao{
		group:   "default",
		table:   "han_admin_group",
		columns: hanAdminGroupColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HanAdminGroupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HanAdminGroupDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HanAdminGroupDao) Columns() HanAdminGroupColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HanAdminGroupDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HanAdminGroupDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HanAdminGroupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
