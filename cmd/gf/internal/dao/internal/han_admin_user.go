// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-05-10 09:06:44
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HanAdminUserDao is the data access object for table han_admin_user.
type HanAdminUserDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns HanAdminUserColumns // columns contains all the column names of Table for convenient usage.
}

// HanAdminUserColumns defines and stores column names for table han_admin_user.
type HanAdminUserColumns struct {
	Id         string // 用户标识
	Username   string // 登陆账号
	Password   string // 登陆密码
	Salt       string // 加密盐值
	Realname   string // 用户名称
	Mobile     string // 手机号码
	Status     string // 状态 0: 未使用 1:正常使用 2:已禁用
	Avatar     string // 头像
	LoginTime  string // 最近一次登陆时间
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	DeleteTime string // 删除时间
}

// hanAdminUserColumns holds the columns for table han_admin_user.
var hanAdminUserColumns = HanAdminUserColumns{
	Id:         "id",
	Username:   "username",
	Password:   "password",
	Salt:       "salt",
	Realname:   "realname",
	Mobile:     "mobile",
	Status:     "status",
	Avatar:     "avatar",
	LoginTime:  "login_time",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	DeleteTime: "delete_time",
}

// NewHanAdminUserDao creates and returns a new DAO object for table data access.
func NewHanAdminUserDao() *HanAdminUserDao {
	return &HanAdminUserDao{
		group:   "default",
		table:   "han_admin_user",
		columns: hanAdminUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HanAdminUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HanAdminUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HanAdminUserDao) Columns() HanAdminUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HanAdminUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HanAdminUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HanAdminUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
