// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-05-10 09:06:44
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// HanAdminRole is the golang structure of table han_admin_role for DAO operations like Where/Data.
type HanAdminRole struct {
	g.Meta     `orm:"table:han_admin_role, do:true"`
	Id         int         `            hashids:"true"  ` // 标识
	Name       interface{} // 角色名称
	Pid        interface{} // 上下级
	CreateTime *gtime.Time // 创建时间
	DeleteTime *gtime.Time // 删除时间
}
