/**
 *
 * @author 半城风雨
 * @since 2021/11/13
 * @File : link
 */
package vo

import "easygoadmin/app/model"

// 友链信息
type LinkInfoVo struct {
	model.Link
	TypeName     string `json:"typeName"`     // 友链类型
	FormName     string `json:"formName"`     // 友链形式
	PlatformName string `json:"platformName"` // 投放平台
}
