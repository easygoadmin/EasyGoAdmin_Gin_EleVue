/**
 *
 * @author 半城风雨
 * @since 2021/11/13
 * @File : member
 */
package vo

import "easygoadmin/app/model"

// 会员信息Vo
type MemberInfoVo struct {
	model.Member
	GenderName string      `json:"genderName"` // 性别
	DeviceName string      `json:"deviceName"` // 设备类型
	SourceName string      `json:"sourceName"` // 会员来源
	City       interface{} `json:"city"`       // 省市区
	CityName   string      `json:"cityName"`   // 城市名称
}
