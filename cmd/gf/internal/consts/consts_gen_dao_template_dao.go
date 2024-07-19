// Copyright GoFrame gf Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package consts

const TemplateGenDaoIndexContent = `
// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package {TplPackageName}

import (
	"{TplImportPrefix}/internal"
)

// internal{TplTableNameCamelCase}Dao is internal type for wrapping internal DAO implements.
type internal{TplTableNameCamelCase}Dao = *internal.{TplTableNameCamelCase}Dao

// {TplTableNameCamelLowerCase}Dao is the data access object for table {TplTableName}.
// You can define custom methods on it to extend its functionality as you wish.
type {TplTableNameCamelLowerCase}Dao struct {
	internal{TplTableNameCamelCase}Dao
}

var (
	// {TplTableNameCamelCase} is globally public accessible object for table {TplTableName} operations.
	{TplTableNameCamelCase} = {TplTableNameCamelLowerCase}Dao{
		internal.New{TplTableNameCamelCase}Dao(),
	}
)

// Fill with you ideas below.

`

// 自定义事务模板
const TemplateGenDaoTxLevelTemplate = `
// ==========================================================================
// The default transaction level switch
// ==========================================================================

package internal

import (
	"context"
	"regexp"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type TransactionLevel string

const (
	ReadUncommitted = "SET SESSION TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;"
	ReadCommitted   = "SET SESSION TRANSACTION ISOLATION LEVEL READ COMMITTED;"
	RepeatableRead  = "SET SESSION TRANSACTION ISOLATION LEVEL REPEATABLE READ;"
	Serializable    = "SET SESSION TRANSACTION ISOLATION LEVEL SERIALIZABLE;"
	MySQL           = "SELECT @@TRANSACTION_ISOLATION AS isolation;"
	QueryVersion    = "SELECT VERSION() AS version"
)

const (
	ReadUncommittedLevel TransactionLevel = "READ-UNCOMMITTED"
	ReadCommittedLevel   TransactionLevel = "READ-COMMITTED"
	RepeatableReadLevel  TransactionLevel = "REPEATABLE-READ"
	SerializableLevel    TransactionLevel = "SERIALIZABLE"
)

var (
	defaultFunc        = func() {}
	mysqlVersionReg, _ = regexp.Compile("\\d+.\\d+.\\d+")
	transactionLevels  = map[TransactionLevel]string{
		ReadUncommittedLevel: ReadUncommitted,
		ReadCommittedLevel:   ReadCommitted,
		RepeatableReadLevel:  RepeatableRead,
		SerializableLevel:    Serializable,
	}
)

func SetTransactionLevel(ctx context.Context, level TransactionLevel) (db gdb.DB, switchFunc func()) {
	db = g.DB()
	switchFunc = defaultFunc
	result, err := db.Query(ctx, QueryVersion)
	if err != nil || result.IsEmpty() {
		return
	}
	version := result[0]["version"].String()
	if version == "" || !mysqlVersionReg.MatchString(version) {
		return
	}
	// mysql版本比较
	result, err = db.Query(ctx, MySQL)
	if err != nil || result.IsEmpty() || TransactionLevel(result[0]["isolation"].String()) == level {
		return
	}
	if _, ok := transactionLevels[level]; !ok {
		return
	}
	_, err = db.Exec(ctx, transactionLevels[level])
	if err != nil {
		panic(err)
	}
	switchFunc = func() {
		_, _ = db.Exec(ctx, transactionLevels[TransactionLevel(result[0]["isolation"].String())])
	}
	return
}
`

const TemplateGenDaoInternalContent = `
// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. {TplCreatedAtDatetimeStr}
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// {TplTableNameCamelCase}Dao is the data access object for table {TplTableName}.
type {TplTableNameCamelCase}Dao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns {TplTableNameCamelCase}Columns // columns contains all the column names of Table for convenient usage.
}

// {TplTableNameCamelCase}Columns defines and stores column names for table {TplTableName}.
type {TplTableNameCamelCase}Columns struct {
	{TplColumnDefine}
}

// {TplTableNameCamelLowerCase}Columns holds the columns for table {TplTableName}.
var {TplTableNameCamelLowerCase}Columns = {TplTableNameCamelCase}Columns{
	{TplColumnNames}
}


// As returns the list of field aliases in table {TplTableName}.
func (c *{TplTableNameCamelCase}Columns) As(as string) {TplTableNameCamelCase}Columns {
	return {TplTableNameCamelCase}Columns{
		{TplVarColumnNameAlias}
	}
}


// New{TplTableNameCamelCase}Dao creates and returns a new DAO object for table data access.
func New{TplTableNameCamelCase}Dao() *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{
		group:   "{TplGroupName}",
		table:   "{TplTableName}",
		columns: {TplTableNameCamelLowerCase}Columns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *{TplTableNameCamelCase}Dao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *{TplTableNameCamelCase}Dao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *{TplTableNameCamelCase}Dao) Columns() {TplTableNameCamelCase}Columns {
	return dao.columns
}

// ColumnAs returns the alias of the current DAO field
func (dao *{TplTableNameCamelCase}Dao) ColumnAs(as string) {TplTableNameCamelCase}Columns {
	return dao.columns.As(as)
}

// Model returns the model class based on the db interface
func (dao *{TplTableNameCamelCase}Dao) Model(ctx context.Context, db gdb.DB) *gdb.Model {
	return db.Model(dao.table).Safe().Ctx(ctx)
}


// ReadUncommittedModel The current session switches to the Read Uncommitted Transaction level
func (dao *{TplTableNameCamelCase}Dao) ReadUncommittedModel(ctx context.Context) (m *gdb.Model, switchFunc func()) {
	db, switchFunc := SetTransactionLevel(ctx, ReadUncommittedLevel)
	return dao.Model(ctx, db).Safe().Ctx(ctx), switchFunc
}

// ReadCommittedModel The current session switches to the Read Committed Transaction level
func (dao *{TplTableNameCamelCase}Dao) ReadCommittedModel(ctx context.Context) (m *gdb.Model, switchFunc func()) {
	db, switchFunc := SetTransactionLevel(ctx, ReadCommittedLevel)
	return dao.Model(ctx, db).Safe().Ctx(ctx), switchFunc
}

// RepeatableReadModel The current session switches to the Repeatable Read Transaction level
func (dao *{TplTableNameCamelCase}Dao) RepeatableReadModel(ctx context.Context) (m *gdb.Model, switchFunc func()) {
	db, switchFunc := SetTransactionLevel(ctx, RepeatableReadLevel)
	return dao.Model(ctx, db).Safe().Ctx(ctx), switchFunc
}

// SerializableModel The current session switches to the Serializable Transaction level
func (dao *{TplTableNameCamelCase}Dao) SerializableModel(ctx context.Context) (m *gdb.Model, switchFunc func()) {
	db, switchFunc := SetTransactionLevel(ctx, SerializableLevel)
	return dao.Model(ctx, db).Safe().Ctx(ctx), switchFunc
}


// Group returns the configuration group name of database of current dao.
func (dao *{TplTableNameCamelCase}Dao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *{TplTableNameCamelCase}Dao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *{TplTableNameCamelCase}Dao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
`