/**
 *
 * @author 半城风雨
 * @since 2021/11/15
 * @File : config_web
 */
package controller

import (
	"easygoadmin/app/model"
	"easygoadmin/utils"
	"easygoadmin/utils/common"
	"easygoadmin/utils/gconv"
	"easygoadmin/utils/gstr"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"regexp"
)

// 控制器管理对象
var ConfigWeb = new(configWeb)

type configWeb struct{}

func (c *configWeb) Index(ctx *gin.Context) {
	// 获取配置列表
	configData := make([]model.Config, 0)
	utils.XormDb.Where("mark=1").Find(&configData)
	configList := make([]map[string]interface{}, 0)
	for _, v := range configData {
		item := make(map[string]interface{})
		item["config_id"] = v.Id
		item["config_name"] = v.Name

		// 查询配置项列表
		itemData := make([]model.ConfigData, 0)
		utils.XormDb.Where("config_id=? and status=1 and mark=1", v.Id).OrderBy("sort asc").Find(&itemData)
		itemList := make([]map[string]interface{}, 0)
		for _, v := range itemData {
			item := make(map[string]interface{})
			item["id"] = v.Id
			item["title"] = v.Title
			item["code"] = v.Code
			item["value"] = v.Value
			item["type"] = v.Type

			if v.Type == "checkbox" {
				// 复选框
				re := regexp.MustCompile(`\r?\n`)
				list := gstr.Split(re.ReplaceAllString(v.Options, "|"), "|")
				optionsList := make(map[string]string)
				for _, val := range list {
					re2 := regexp.MustCompile(`:|：|\s+`)
					item := gstr.Split(re2.ReplaceAllString(val, "|"), "|")
					optionsList[item[0]] = item[1]
				}
				// 选择项
				item["optionsList"] = optionsList
				// 选择值
				item["value"] = gstr.Split(v.Value, ",")

			} else if v.Type == "radio" {
				// 单选框
				re := regexp.MustCompile(`\r?\n`)
				list := gstr.Split(re.ReplaceAllString(v.Options, "|"), "|")
				optionsList := make(map[string]string)
				for _, v := range list {
					re2 := regexp.MustCompile(`:|：|\s+`)
					item := gstr.Split(re2.ReplaceAllString(v, "|"), "|")
					optionsList[item[0]] = item[1]
				}
				item["optionsList"] = optionsList

			} else if v.Type == "select" {
				// 下拉选择框
				re := regexp.MustCompile(`\r?\n`)
				list := gstr.Split(re.ReplaceAllString(v.Options, "|"), "|")
				optionsList := make(map[string]string)
				for _, v := range list {
					re2 := regexp.MustCompile(`:|：|\s+`)
					item := gstr.Split(re2.ReplaceAllString(v, "|"), "|")
					optionsList[item[0]] = item[1]
				}
				item["optionsList"] = optionsList
			} else if v.Type == "image" {
				// 单图片
				item["value"] = utils.GetImageUrl(v.Value)
			} else if v.Type == "images" {
				// 多图片
				list := gstr.Split(v.Value, ",")
				itemList := make([]string, 0)
				for _, v := range list {
					// 图片地址
					item := utils.GetImageUrl(v)
					itemList = append(itemList, item)
				}
				item["value"] = itemList
			}
			itemList = append(itemList, item)
		}
		item["itemList"] = itemList
		// 加入数组
		configList = append(configList, item)
	}

	// 返回结果
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code: 0,
		Msg:  "查询成功",
		Data: configList,
	})
}

func (c *configWeb) Save(ctx *gin.Context) {
	// key：string类型，value：interface{}  类型能存任何数据类型
	var jsonObj map[string]interface{}
	data, _ := ctx.GetRawData()
	json.Unmarshal(data, &jsonObj)
	for key, val := range jsonObj {
		// 数组处理
		if reflect.ValueOf(val).Kind() == reflect.Slice {
			if reflect.TypeOf(val).String() == "[]interface {}" {
				// 初始化URL数组
				item := make([]string, 0)
				for _, v := range val.([]interface{}) {
					value := gconv.String(v)
					// 判断是否http(s)开头
					if gstr.SubStr(value, 0, 4) == "http" ||
						gstr.SubStr(value, 0, 5) == "https" {
						// 图片地址
						if gstr.Contains(value, utils.ImageUrl()) {
							url, _ := utils.SaveImage(value, "config")
							item = append(item, url)
						}
					} else {
						// 复选框处理
						item = append(item, value)
					}
				}
				// 逗号拼接
				val = gstr.Join(item, ",")
			}
		} else {
			// 图片处理
			if gstr.Contains(gconv.String(val), "http://") ||
				gstr.Contains(gconv.String(val), "https://") {
				// 图片地址
				if gstr.Contains(gconv.String(val), utils.ImageUrl()) {
					val, _ = utils.SaveImage(gconv.String(val), "config")
				}
			}
		}

		// 查询记录
		info := &model.ConfigData{}
		has, err := utils.XormDb.Where("code", key).Get(&info) //dao.ConfigData.FindOne("code=?", key)
		if err != nil || !has {
			continue
		}

		// 更新记录
		itemVal := &model.ConfigData{}
		itemVal.Value = gconv.String(val)
		_, err2 := utils.XormDb.Where("code=?", key).Update(&itemVal)
		if err2 != nil {
			continue
		}
	}

	// 返回结果
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code: 0,
		Msg:  "保存成功",
	})
}
