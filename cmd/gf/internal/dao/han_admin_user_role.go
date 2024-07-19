// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/gogf/gf/cmd/gf/v2/internal/dao/internal"
)

// internalHanAdminUserRoleDao is internal type for wrapping internal DAO implements.
type internalHanAdminUserRoleDao = *internal.HanAdminUserRoleDao

// hanAdminUserRoleDao is the data access object for table han_admin_user_role.
// You can define custom methods on it to extend its functionality as you wish.
type hanAdminUserRoleDao struct {
	internalHanAdminUserRoleDao
}

var (
	// HanAdminUserRole is globally public accessible object for table han_admin_user_role operations.
	HanAdminUserRole = hanAdminUserRoleDao{
		internal.NewHanAdminUserRoleDao(),
	}
)

// Fill with you ideas below.
