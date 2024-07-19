// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/gogf/gf/cmd/gf/v2/internal/dao/internal"
)

// internalHanAdminRoleMutexDao is internal type for wrapping internal DAO implements.
type internalHanAdminRoleMutexDao = *internal.HanAdminRoleMutexDao

// hanAdminRoleMutexDao is the data access object for table han_admin_role_mutex.
// You can define custom methods on it to extend its functionality as you wish.
type hanAdminRoleMutexDao struct {
	internalHanAdminRoleMutexDao
}

var (
	// HanAdminRoleMutex is globally public accessible object for table han_admin_role_mutex operations.
	HanAdminRoleMutex = hanAdminRoleMutexDao{
		internal.NewHanAdminRoleMutexDao(),
	}
)

// Fill with you ideas below.