// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-05-10 09:06:44
// =================================================================================

package entity

// HanAdminRoleMutex is the golang structure for table han_admin_role_mutex.
type HanAdminRoleMutex struct {
	Id  int `json:"id,omitempty"  hashids:"true" orm:"id"               description:"标识"  ` // 标识
	Rid int `json:"rid,omitempty" orm:"rid"      description:"角色id"    `                    // 角色id
	Mid int `json:"mid,omitempty" orm:"mid"      description:"互斥角色"  `                      // 互斥角色
}
