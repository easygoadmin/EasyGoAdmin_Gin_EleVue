// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 EasyGoAdmin深圳研发中心
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: 半城风雨 <easygoadmin@163.com>
// +----------------------------------------------------------------------
// | 免责声明:
// | 本软件框架禁止任何单位和个人用于任何违法、侵害他人合法利益等恶意的行为，禁止用于任何违
// | 反我国法律法规的一切平台研发，任何单位和个人使用本软件框架用于产品研发而产生的任何意外
// | 、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、附带
// | 或衍生的损失等)，本团队不承担任何法律责任。本软件框架只能用于公司和个人内部的法律所允
// | 许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

/**
 * 公共函数库
 * @author 半城风雨
 * @since 2021/3/2
 * @File : common
 */
package common

type BunissType int

//业务类型
const (
	BOther BunissType = 0 //0其它
	BAdd   BunissType = 1 //1新增
	BEdit  BunissType = 2 //2修改
	BDel   BunissType = 3 //3删除
)

type JsonResult struct {
	Code  int         `json:"code"`  // 响应编码：0成功 401请登录 403无权限 500错误
	Msg   string      `json:"msg"`   // 消息提示语
	Data  interface{} `json:"data"`  // 数据对象
	Count int64       `json:"count"` // 记录总数
	Btype BunissType  `json:"btype"` // 业务类型
}

type CaptchaRes struct {
	Code  int         `json:"code"`  //响应编码 0 成功 500 错误 403 无权限
	Msg   string      `json:"msg"`   //消息
	Data  interface{} `json:"data"`  //数据内容
	IdKey string      `json:"idkey"` //验证码ID
}

// 部门类型
var DEPT_TYPE_LIST = map[int]string{
	1: "公司",
	2: "子公司",
	3: "部门",
	4: "小组",
}

// 菜单类型
var MENU_TYPE_LIST = map[int]string{
	0: "菜单",
	1: "节点",
}

// 城市等级
var CITY_LEVEL = map[int]string{
	1: "省份",
	2: "城市",
	3: "县区",
	4: "街道",
}

// 配置项类型
var CONFIG_DATA_TYPE_LIST = map[string]string{
	"text":     "单行文本",
	"textarea": "多行文本",
	"ueditor":  "富文本编辑器",
	"date":     "日期",
	"datetime": "时间",
	"number":   "数字",
	"select":   "下拉框",
	"radio":    "单选框",
	"checkbox": "复选框",
	"image":    "单张图片",
	"images":   "多张图片",
	"password": "密码",
	"icon":     "字体图标",
	"file":     "单个文件",
	"files":    "多个文件",
	"hidden":   "隐藏",
	"readonly": "只读文本",
}

// 友链类型
var LINK_TYPE_LIST = map[int]string{
	1: "友情链接",
	2: "合作伙伴",
}

// 友链形式
var LINK_FORM_LIST = map[int]string{
	1: "文字链接",
	2: "图片链接",
}

// 友链平台
var LINK_PLATFORM_LIST = map[int]string{
	1: "PC站",
	2: "WAP站",
	3: "小程序",
	4: "APP应用",
}

// 站点类型
var ITEM_TYPE_LIST = map[int]string{
	1: "国内站点",
	2: "国外站点",
	3: "其他站点",
}

// 广告位所属平台
var ADSORT_PLATFORM_LIST = map[int]string{
	1: "PC站",
	2: "WAP站",
	3: "小程序",
	4: "APP应用",
}

// 广告类型
var AD_TYPE_LIST = map[int]string{
	1: "图片",
	2: "文字",
	3: "视频",
	4: "其他",
}

// 通知来源
var NOTICE_SOURCE_LIST = map[int]string{
	1: "内部通知",
	2: "外部通知",
}

// 会员设备类型
var MEMBER_DEVICE_LIST = map[int]string{
	1: "苹果",
	2: "安卓",
	3: "WAP站",
	4: "PC站",
	5: "后台添加",
}

// 会员来源
var MEMBER_SOURCE_LIST = map[int]string{
	1: "注册会员",
	2: "马甲会员",
}
