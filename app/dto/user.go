// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本团队对该软件框架产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架禁止任何单位和个人、组织用于
// | 任何违法、侵害他人合法利益等恶意的行为，如有发现违规、违法的犯罪行为，本团队将无条件配
// | 合公安机关调查取证同时保留一切以法律手段起诉的权利，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

/**
 * 用户Dto
 * @author 半城风雨
 * @since 2021/11/11
 * @File : user
 */
package dto

// 用户分页查询条件
type UserPageReq struct {
	Username string `form:"username"` // 用户名
	Gender   int    `form:gender`     // 性别
	Page     int    `form:"page"`     // 页码
	Limit    int    `form:"limit"`    // 每页数
}

// 添加用户
type UserAddReq struct {
	Realname   string   `form:"realname" binding:"required"`
	Nickname   string   `form:"nickname" binding:"required"`
	Gender     int      `form:"gender" binding:"required"`
	Avatar     string   `form:"avatar" binding:"required"`
	Mobile     string   `form:"mobile" binding:"required"`
	Email      string   `form:"email" binding:"required"`
	Birthday   int64    `form:"birthday" binding:"required"`
	DeptId     int      `form:"dept_id" binding:"required"`
	LevelId    int      `form:"level_id" binding:"required"`
	PositionId int      `form:"position_id" binding:"required"`
	City       []string `form:"city" binding:"required"` // 省市区
	Address    string   `form:"address"`
	Username   string   `form:"username" binding:"required"`
	Password   string   `form:"password"`
	Intro      string   `form:"intro"`
	Status     int      `form:"status" binding:"required"`
	Note       string   `form:"note"`
	Sort       int      `form:"sort"`
	RoleIds    []int    `form:"roleIds"` // 用户角色
}

// 更新用户
type UserUpdateReq struct {
	Id         int      `form:"id" binding:"required"`
	Realname   string   `form:"realname" binding:"required"`
	Nickname   string   `form:"nickname" binding:"required"`
	Gender     int      `form:"gender" binding:"required"`
	Avatar     string   `form:"avatar" binding:"required"`
	Mobile     string   `form:"mobile" binding:"required"`
	Email      string   `form:"email" binding:"required"`
	Birthday   int64    `form:"birthday" binding:"required"`
	DeptId     int      `form:"dept_id" binding:"required"`
	LevelId    int      `form:"level_id" binding:"required"`
	PositionId int      `form:"position_id" binding:"required"`
	City       []string `form:"city" binding:"required"` // 省市区
	Address    string   `form:"address"`
	Username   string   `form:"username" binding:"required"`
	Password   string   `form:"password"`
	Intro      string   `form:"intro"`
	Status     int      `form:"status" binding:"required"`
	Note       string   `form:"note"`
	Sort       int      `form:"sort"`
	RoleIds    []int    `form:"roleIds"` // 用户角色
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
