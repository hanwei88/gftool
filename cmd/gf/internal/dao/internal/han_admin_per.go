// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-05-10 09:06:44
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HanAdminPerDao is the data access object for table han_admin_per.
type HanAdminPerDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns HanAdminPerColumns // columns contains all the column names of Table for convenient usage.
}

// HanAdminPerColumns defines and stores column names for table han_admin_per.
type HanAdminPerColumns struct {
	Id         string //
	Title      string // 名称
	Type       string // 类型 0: 分组 2: 页面 3: 按钮(操作)
	Pid        string // 上级
	Content    string // 补充内容
	CreateTime string // 创建时间
	DeleteTime string // 删除时间
}

// hanAdminPerColumns holds the columns for table han_admin_per.
var hanAdminPerColumns = HanAdminPerColumns{
	Id:         "id",
	Title:      "title",
	Type:       "type",
	Pid:        "pid",
	Content:    "content",
	CreateTime: "create_time",
	DeleteTime: "delete_time",
}

// NewHanAdminPerDao creates and returns a new DAO object for table data access.
func NewHanAdminPerDao() *HanAdminPerDao {
	return &HanAdminPerDao{
		group:   "default",
		table:   "han_admin_per",
		columns: hanAdminPerColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HanAdminPerDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HanAdminPerDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HanAdminPerDao) Columns() HanAdminPerColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HanAdminPerDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HanAdminPerDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HanAdminPerDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
