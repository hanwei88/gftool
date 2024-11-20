// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-05-10 09:06:44
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// HanAdminPerApi is the golang structure for table han_admin_per_api.
type HanAdminPerApi struct {
	Id         int         `json:"id,omitempty"          hashids:"true"    orm:"id"               description:""  ` //
	PerId      int         `json:"per_id,omitempty"      orm:"per_id"      description:""          `                //
	Path       string      `json:"path,omitempty"        orm:"path"        description:""          `                //
	CreateTime *gtime.Time `json:"create_time,omitempty" orm:"create_time" description:"创建时间"  `                    // 创建时间
}
