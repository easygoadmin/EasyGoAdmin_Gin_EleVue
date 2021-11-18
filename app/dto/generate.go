/**
 *
 * @author 半城风雨
 * @since 2021/11/15
 * @File : generate
 */
package dto

// 分页查询条件
type GeneratePageReq struct {
	Name    string `form:"name"`    // 表名称
	Comment string `form:"comment"` // 表描述
	Page    int    `form:"page"`    // 页码
	Limit   int    `form:"limit"`   // 每页数
}

// 生成文件
type GenerateFileReq struct {
	Name    string `form:"name"`    // 表名称
	Comment string `form:"comment"` // 表描述
}
