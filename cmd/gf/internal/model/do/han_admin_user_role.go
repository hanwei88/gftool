// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-05-10 09:06:44
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// HanAdminUserRole is the golang structure of table han_admin_user_role for DAO operations like Where/Data.
type HanAdminUserRole struct {
	g.Meta     `orm:"table:han_admin_user_role, do:true"`
	Id         int         `            hashids:"true"  ` // 用户关系标识
	Uid        interface{} // 用户id
	Rid        interface{} // 角色id
	CreateTime *gtime.Time // 创建时间
}
