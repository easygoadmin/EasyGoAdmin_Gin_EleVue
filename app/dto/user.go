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
