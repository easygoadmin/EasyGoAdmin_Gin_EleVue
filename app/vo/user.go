/**
 *
 * @author 半城风雨
 * @since 2021/11/12
 * @File : user
 */
package vo

import "easygoadmin/app/model"

// 用户信息Vo
type UserInfoVo struct {
	model.User
	GenderName   string      `json:"genderName"`   // 性别
	LevelName    string      `json:"levelName"`    // 职级
	PositionName string      `json:"positionName"` // 岗位
	DeptName     string      `json:"deptName"`     // 部门
	RoleIds      interface{} `json:"roleIds"`      // 角色ID
	RoleList     interface{} `json:"roleList"`     // 角色列表
	City         interface{} `json:"city"`         // 省市区
}
