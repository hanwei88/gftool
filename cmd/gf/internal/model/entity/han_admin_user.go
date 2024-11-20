// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-05-10 09:06:44
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// HanAdminUser is the golang structure for table han_admin_user.
type HanAdminUser struct {
	Id         int         `json:"id,omitempty"          hashids:"true"    orm:"id"                                         description:"用户标识"  ` // 用户标识
	Username   string      `json:"username,omitempty"    orm:"username"    description:"登陆账号"                            `                        // 登陆账号
	Password   string      `json:"-"                     orm:"password"    description:"登陆密码"                            `                        // 登陆密码
	Salt       string      `json:"-"                     orm:"salt"        description:"加密盐值"                            `                        // 加密盐值
	Realname   string      `json:"realname,omitempty"    orm:"realname"    description:"用户名称"                            `                        // 用户名称
	Mobile     string      `json:"mobile,omitempty"      orm:"mobile"      description:"手机号码"                            `                        // 手机号码
	Status     uint8       `json:"status,omitempty"      orm:"status"      description:"状态 0: 未使用 1:正常使用 2:已禁用"  `                                // 状态 0: 未使用 1:正常使用 2:已禁用
	Avatar     string      `json:"avatar,omitempty"      orm:"avatar"      description:"头像"                                `                      // 头像
	LoginTime  *gtime.Time `json:"login_time,omitempty"  orm:"login_time"  description:"最近一次登陆时间"                    `                            // 最近一次登陆时间
	CreateTime *gtime.Time `json:"create_time,omitempty" orm:"create_time" description:"创建时间"                            `                        // 创建时间
	UpdateTime *gtime.Time `json:"update_time,omitempty" orm:"update_time" description:"更新时间"                            `                        // 更新时间
	DeleteTime *gtime.Time `json:"-"                     orm:"delete_time" description:"删除时间"                            `                        // 删除时间
}
