/**
 *
 * @author 半城风雨
 * @since 2021/11/15
 * @File : city
 */
package vo

import "easygoadmin/app/model"

type CityInfoVo struct {
	model.City
	HaveChild bool `json:"haveChild"`
}
