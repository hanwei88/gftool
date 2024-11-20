// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-05-10 09:06:44
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// HanAdminGroup is the golang structure for table han_admin_group.
type HanAdminGroup struct {
	Id         int         `json:"id,omitempty"          hashids:"true"    orm:"id"                                   description:"分组标识"  ` // 分组标识
	Name       string      `json:"name,omitempty"        orm:"name"        description:"分组名称"                      `                        // 分组名称
	Type       int8        `json:"type,omitempty"        orm:"type"        description:"类型 0：分组或者部门 1：职位"  `                                // 类型 0：分组或者部门 1：职位
	Pid        int         `json:"pid,omitempty"         orm:"pid"         description:"上下级"                        `                       // 上下级
	CreateTime *gtime.Time `json:"create_time,omitempty" orm:"create_time" description:"创建时间"                      `                        // 创建时间
	DeleteTime *gtime.Time `json:"-"                     orm:"delete_time" description:"删除时间"                      `                        // 删除时间
}
