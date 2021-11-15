/**
 *
 * @author 半城风雨
 * @since 2021/11/15
 * @File : ad
 */
package vo

import "easygoadmin/app/model"

// 广告信息Vo
type AdInfoVo struct {
	model.Ad
	TypeName   string `json:"typeName"`   // 广告类型
	AdSortDesc string `json:"adSortDesc"` // 广告位描述
}
