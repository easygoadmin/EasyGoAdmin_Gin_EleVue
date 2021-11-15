/**
 *
 * @author 半城风雨
 * @since 2021/11/12
 * @File : oper_log
 */
package dto

// 分页查询条件
type OperLogPageReq struct {
	Username string `form:"username"` // 操作账号
	Model    string `form:"model"`    // 操作模块
	Page     int    `form:"page"`     // 页码
	Limit    int    `form:"limit"`    // 每页数
}
