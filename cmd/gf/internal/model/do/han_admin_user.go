// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-05-10 09:06:44
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// HanAdminUser is the golang structure of table han_admin_user for DAO operations like Where/Data.
type HanAdminUser struct {
	g.Meta     `orm:"table:han_admin_user, do:true"`
	Id         int         `                                      hashids:"true"  ` // 用户标识
	Username   interface{} // 登陆账号
	Password   interface{} // 登陆密码
	Salt       interface{} // 加密盐值
	Realname   interface{} // 用户名称
	Mobile     interface{} // 手机号码
	Status     interface{} // 状态 0: 未使用 1:正常使用 2:已禁用
	Avatar     interface{} // 头像
	LoginTime  *gtime.Time // 最近一次登陆时间
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	DeleteTime *gtime.Time // 删除时间
}
