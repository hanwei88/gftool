// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-05-10 09:06:44
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// HanAdminPer is the golang structure of table han_admin_per for DAO operations like Where/Data.
type HanAdminPer struct {
	g.Meta     `orm:"table:han_admin_per, do:true"`
	Id         int         `                                      hashids:"true"  ` //
	Title      interface{} // 名称
	Type       interface{} // 类型 0: 分组 2: 页面 3: 按钮(操作)
	Pid        interface{} // 上级
	Content    interface{} // 补充内容
	CreateTime *gtime.Time // 创建时间
	DeleteTime *gtime.Time // 删除时间
}
