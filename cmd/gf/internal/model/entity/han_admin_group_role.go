// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-05-10 09:06:44
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// HanAdminGroupRole is the golang structure for table han_admin_group_role.
type HanAdminGroupRole struct {
	Id         int         `json:"id,omitempty"          hashids:"true"    orm:"id"               description:"分组标识"  ` // 分组标识
	Gid        int         `json:"gid,omitempty"         orm:"gid"         description:"分组id"    `                      // 分组id
	Rid        int         `json:"rid,omitempty"         orm:"rid"         description:"角色id"    `                      // 角色id
	CreateTime *gtime.Time `json:"create_time,omitempty" orm:"create_time" description:"创建时间"  `                        // 创建时间
}
