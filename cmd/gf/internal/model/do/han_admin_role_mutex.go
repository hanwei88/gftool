// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-05-10 09:06:44
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// HanAdminRoleMutex is the golang structure of table han_admin_role_mutex for DAO operations like Where/Data.
type HanAdminRoleMutex struct {
	g.Meta `orm:"table:han_admin_role_mutex, do:true"`
	Id     int         `            hashids:"true"  ` // 标识
	Rid    interface{} // 角色id
	Mid    interface{} // 互斥角色
}
