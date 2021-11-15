/**
 *
 * @author 半城风雨
 * @since 2021/11/13
 * @File : item
 */
package vo

import "easygoadmin/app/model"

// 站点信息Vo
type ItemInfoVo struct {
	model.Item
	TypeName string `json:"typeName"` // 站点类型
}
