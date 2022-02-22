// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 EasyGoAdmin深圳研发中心
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: 半城风雨 <easygoadmin@163.com>
// +----------------------------------------------------------------------
// | 免责声明:
// | 本软件框架禁止任何单位和个人用于任何违法、侵害他人合法利益等恶意的行为，禁止用于任何违
// | 反我国法律法规的一切平台研发，任何单位和个人使用本软件框架用于产品研发而产生的任何意外
// | 、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、附带
// | 或衍生的损失等)，本团队不承担任何法律责任。本软件框架只能用于公司和个人内部的法律所允
// | 许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

/**
 * 用户Vo
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

// 个人信息Vo
type ProfileInfoVo struct {
	Realname       string        `json:"realname"`       // 真实姓名
	Nickname       string        `json:"nickname"`       // 昵称
	Gender         int           `json:"gender"`         // 性别:1男 2女 3保密
	Avatar         string        `json:"avatar"`         // 头像
	Mobile         string        `json:"mobile"`         // 手机号码
	Email          string        `json:"email"`          // 邮箱地址
	City           []string      `json:"city"`           // 省市区
	Address        string        `json:"address"`        // 详细地址
	Intro          string        `json:"intro"`          // 个人简介
	Roles          []interface{} `json:"roles"`          // 用户角色
	Authorities    []interface{} `json:"authorities"`    // 用户权限
	PermissionList []string      `json:"permissionList"` // 权限列表
}
