/**
 *
 * @author 半城风雨
 * @since 2021/11/13
 * @File : notice
 */
package vo

import "easygoadmin/app/model"

// 通知公告Vo
type NoticeInfoVo struct {
	model.Notice
	SourceName string `json:"sourceName"` // 通知来源
}
