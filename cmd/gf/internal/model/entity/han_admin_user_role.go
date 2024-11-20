// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-05-10 09:06:44
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// HanAdminUserRole is the golang structure for table han_admin_user_role.
type HanAdminUserRole struct {
	Id         int         `json:"id,omitempty"          hashids:"true"    orm:"id"               description:"用户关系标识"  ` // 用户关系标识
	Uid        int         `json:"uid,omitempty"         orm:"uid"         description:"用户id"    `                        // 用户id
	Rid        int         `json:"rid,omitempty"         orm:"rid"         description:"角色id"    `                        // 角色id
	CreateTime *gtime.Time `json:"create_time,omitempty" orm:"create_time" description:"创建时间"  `                          // 创建时间
}
