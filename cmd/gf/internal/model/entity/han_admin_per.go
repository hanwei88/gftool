// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-05-10 09:06:44
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// HanAdminPer is the golang structure for table han_admin_per.
type HanAdminPer struct {
	Id         int         `json:"id,omitempty"          hashids:"true"    orm:"id"                                         description:""  ` //
	Title      string      `json:"title,omitempty"       orm:"title"       description:"名称"                                `                  // 名称
	Type       int8        `json:"type,omitempty"        orm:"type"        description:"类型 0: 分组 2: 页面 3: 按钮(操作)"  `                          // 类型 0: 分组 2: 页面 3: 按钮(操作)
	Pid        int         `json:"pid,omitempty"         orm:"pid"         description:"上级"                                `                  // 上级
	Content    string      `json:"content,omitempty"     orm:"content"     description:"补充内容"                            `                    // 补充内容
	CreateTime *gtime.Time `json:"create_time,omitempty" orm:"create_time" description:"创建时间"                            `                    // 创建时间
	DeleteTime *gtime.Time `json:"-"                     orm:"delete_time" description:"删除时间"                            `                    // 删除时间
}
