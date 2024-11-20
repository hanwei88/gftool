// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-05-10 09:06:44
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// HanAdminGroup is the golang structure of table han_admin_group for DAO operations like Where/Data.
type HanAdminGroup struct {
	g.Meta     `orm:"table:han_admin_group, do:true"`
	Id         int         `                                hashids:"true"  ` // 分组标识
	Name       interface{} // 分组名称
	Type       interface{} // 类型 0：分组或者部门 1：职位
	Pid        interface{} // 上下级
	CreateTime *gtime.Time // 创建时间
	DeleteTime *gtime.Time // 删除时间
}
