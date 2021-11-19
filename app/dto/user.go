/**
 *
 * @author 半城风雨
 * @since 2021/11/11
 * @File : user
 */
package dto

import "time"

// 用户分页查询条件
type UserPageReq struct {
	Username string `form:"username"` // 用户名
	Gender   int    `form:gender`     // 性别
	Page     int    `form:"page"`     // 页码
	Limit    int    `form:"limit"`    // 每页数
}

// 添加用户
type UserAddReq struct {
	Realname   string    `form:"realname" binding:"required"`
	Nickname   string    `form:"nickname" binding:"required"`
	Gender     int       `form:"gender" binding:"required"`
	Avatar     string    `form:"avatar" binding:"required"`
	Mobile     string    `form:"mobile" binding:"required"`
	Email      string    `form:"email" binding:"required"`
	Birthday   time.Time `form:"birthday" binding:"required"`
	DeptId     int       `form:"dept_id" binding:"required"`
	LevelId    int       `form:"level_id" binding:"required"`
	PositionId int       `form:"position_id" binding:"required"`
	City       []string  `form:"city" binding:"required"` // 省市区
	Address    string    `form:"address"`
	Username   string    `form:"username" binding:"required"`
	Password   string    `form:"password" binding:"required"`
	Intro      string    `form:"intro"`
	Status     int       `form:"status" binding:"required"`
	Note       string    `form:"note"`
	Sort       int       `form:"sort" binding:"required"`
}

// 更新用户
type UserUpdateReq struct {
	Id         int       `form:"id" binding:"required"`
	Realname   string    `form:"realname" binding:"required"`
	Nickname   string    `form:"nickname" binding:"required"`
	Gender     int       `form:"gender" binding:"required"`
	Avatar     string    `form:"avatar" binding:"required"`
	Mobile     string    `form:"mobile" binding:"required"`
	Email      string    `form:"email" binding:"required"`
	Birthday   time.Time `form:"birthday" binding:"required"`
	DeptId     int       `form:"dept_id" binding:"required"`
	LevelId    int       `form:"level_id" binding:"required"`
	PositionId int       `form:"position_id" binding:"required"`
	City       []string  `form:"city" binding:"required"` // 省市区
	Address    string    `form:"address"`
	Username   string    `form:"username" binding:"required"`
	Password   string    `form:"password" binding:"required"`
	Intro      string    `form:"intro"`
	Status     int       `form:"status" binding:"required"`
	Note       string    `form:"note"`
	Sort       int       `form:"sort" binding:"required"`
}

// 设置状态
type UserStatusReq struct {
	Id     int `form:"id" binding:"required"`
	Status int `form:"status"    binding:"required"`
}

// 重置密码
type UserResetPwdReq struct {
	Id int `form:"id" binding:"required"`
}

// 检查用户
type CheckUserReq struct {
	Username string `form:"username" binding:"required"` // 用户名
}

// 更新密码
type UpdatePwd struct {
	OldPassword string `form:"oldPassword"      binding:"required"` // 旧密码
	NewPassword string `form:"newPassword"      binding:"required"` // 新密码
	RePassword  string `form:"rePassword"       binding:"required"` // 确认密码
}

// 用户中心
type UserInfoReq struct {
	Avatar   string `form:"avatar"`                           // 头像
	Realname string `form:"realname"      binding:"required"` // 真实姓名
	Nickname string `form:"nickname"      binding:"required"` // 昵称
	Gender   int    `form:"gender"        binding:"required"` // 性别:1男 2女 3保密
	Mobile   string `form:"mobile"        binding:"required"` // 手机号码
	Email    string `form:"email"         binding:"required"` // 邮箱地址
	Address  string `form:"address"`                          // 详细地址
	Intro    string `form:"intro"`                            // 个人简介
}
