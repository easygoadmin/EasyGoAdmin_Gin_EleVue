/**
 *
 * @author 半城风雨
 * @since 2021/11/12
 * @File : login_log
 */
package dto

// 分页信息查询条件
type LoginLogPageReq struct {
	Username string `form:"username"` // 用户账号
	Page     int    `form:"page"`     // 页码
	Limit    int    `form:"limit"`    // 每页数
}
