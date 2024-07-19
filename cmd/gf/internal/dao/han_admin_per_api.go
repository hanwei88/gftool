// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/gogf/gf/cmd/gf/v2/internal/dao/internal"
)

// internalHanAdminPerApiDao is internal type for wrapping internal DAO implements.
type internalHanAdminPerApiDao = *internal.HanAdminPerApiDao

// hanAdminPerApiDao is the data access object for table han_admin_per_api.
// You can define custom methods on it to extend its functionality as you wish.
type hanAdminPerApiDao struct {
	internalHanAdminPerApiDao
}

var (
	// HanAdminPerApi is globally public accessible object for table han_admin_per_api operations.
	HanAdminPerApi = hanAdminPerApiDao{
		internal.NewHanAdminPerApiDao(),
	}
)

// Fill with you ideas below.
