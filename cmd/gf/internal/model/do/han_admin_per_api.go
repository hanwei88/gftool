// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-05-10 09:06:44
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// HanAdminPerApi is the golang structure of table han_admin_per_api for DAO operations like Where/Data.
type HanAdminPerApi struct {
	g.Meta     `orm:"table:han_admin_per_api, do:true"`
	Id         int         `            hashids:"true"  ` //
	PerId      interface{} //
	Path       interface{} //
	CreateTime *gtime.Time // 创建时间
}
