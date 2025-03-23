package consts;




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

