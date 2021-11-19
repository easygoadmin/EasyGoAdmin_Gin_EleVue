// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2021 EasyGoAdmin深圳研发中心
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
 * 会员Dto
 * @author 半城风雨
 * @since 2021/11/13
 * @File : member
 */
package dto

import "time"

// 分页查询条件
type MemberPageReq struct {
	Username string `form:"username"` // 用户名
	Gender   int    `form:"gender"`   // 性别（1男 2女 3未知）
	Page     int    `form:"page"`     // 页码
	Limit    int    `form:"limit"`    // 每页数
}

// 添加会员
type MemberAddReq struct {
	Username    string    `form:"username,unique" binding:"required"`     // 用户名
	Password    string    `form:"password"		  binding:"required"`   // 登录密码
	MemberLevel int       `form:"member_level"    binding:"required"`     // 会员等级
	Realname    string    `form:"realname"        binding:"required"`     // 真实姓名
	Nickname    string    `form:"nickname"        binding:"required"`     // 用户昵称
	Gender      int       `form:"gender"          binding:"required"`     // 性别（1男 2女 3未知）
	Avatar      string    `form:"avatar"          binding:"required"`     // 用户头像
	Birthday    time.Time `form:"birthday"        binding:"required"`     // 出生日期
	City        []string  `form:"city"		  	  binding:"required"` // 省市区
	Address     string    `form:"address"         binding:"required"`     // 详细地址
	Intro       string    `form:"intro"`                                  // 个人简介
	Signature   string    `form:"signature"`                              // 个性签名
	Device      int       `form:"device"          binding:"required"`     // 设备类型：1苹果 2安卓 3WAP站 4PC站 5后台添加
	Source      int       `form:"source"          binding:"required"`     // 来源：1、APP注册；2、后台添加；
	Status      int       `form:"status"          binding:"required"`     // 是否启用 1、启用  2、停用
}

// 更新会员
type MemberUpdateReq struct {
	Id          int       `form:"id" binding:"required"`
	Username    string    `form:"username,unique" binding:"required"`     // 用户名
	Password    string    `form:"password"		  binding:"required"`   // 登录密码
	MemberLevel int       `form:"member_level"    binding:"required"`     // 会员等级
	Realname    string    `form:"realname"        binding:"required"`     // 真实姓名
	Nickname    string    `form:"nickname"        binding:"required"`     // 用户昵称
	Gender      int       `form:"gender"          binding:"required"`     // 性别（1男 2女 3未知）
	Avatar      string    `form:"avatar"          binding:"required"`     // 用户头像
	Birthday    time.Time `form:"birthday"        binding:"required"`     // 出生日期
	City        []string  `form:"city"		  	  binding:"required"` // 省市区
	Address     string    `form:"address"         binding:"required"`     // 详细地址
	Intro       string    `form:"intro"`                                  // 个人简介
	Signature   string    `form:"signature"`                              // 个性签名
	Device      int       `form:"device"          binding:"required"`     // 设备类型：1苹果 2安卓 3WAP站 4PC站 5后台添加
	Source      int       `form:"source"          binding:"required"`     // 来源：1、APP注册；2、后台添加；
	Status      int       `form:"status"          binding:"required"`     // 是否启用 1、启用  2、停用
}

// 设置状态
type MemberStatusReq struct {
	Id     int `form:"id"        binding:"required"`
	Status int `form:"status"    binding:"required"`
}
