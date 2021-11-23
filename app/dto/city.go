// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2021 EasyGoAdmin深圳研发中心
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
 * 城市Dto
 * @author 半城风雨
 * @since 2021/11/13
 * @File : city
 */
package dto

// 列表查询条件
type CityQueryReq struct {
	Name string `form:"name"` // 城市名称
	Pid  int    `form:"pid"`  // 上级ID
}

// 添加城市
type CityAddReq struct {
	Pid      int    `form:"pid"`                            // 父级编号
	Level    int    `form:"level"       binding:"required"` // 城市级别：1省 2市 3区
	Name     string `form:"name"        binding:"required"` // 城市名称
	Citycode string `form:"citycode"`                       // 城市编号（区号）
	PAdcode  string `form:"p_adcode"`                       // 父级地理编号
	Adcode   string `form:"adcode"`                         // 地理编号
	Lng      string `form:"lng"`                            // 城市坐标中心点经度（* 1e6）：如果是中国，此值是 1e7
	Lat      string `form:"lat"`                            // 城市坐标中心点纬度（* 1e6）
	Sort     int    `form:"sort"`                           // 排序号
}

// 编辑城市
type CityUpdateReq struct {
	Id       int    `form:"id" binding:"required"`          // 主键ID
	Pid      int    `form:"pid"`                            // 父级编号
	Level    int    `form:"level"       binding:"required"` // 城市级别：1省 2市 3区
	Name     string `form:"name"        binding:"required"` // 城市名称
	Citycode string `form:"citycode"`                       // 城市编号（区号）
	PAdcode  string `form:"p_adcode"`                       // 父级地理编号
	Adcode   string `form:"adcode"`                         // 地理编号
	Lng      string `form:"lng"`                            // 城市坐标中心点经度（* 1e6）：如果是中国，此值是 1e7
	Lat      string `form:"lat"`                            // 城市坐标中心点纬度（* 1e6）
	Sort     int    `form:"sort"`                           // 排序号
}

// 获取子级城市
type CityChildReq struct {
	CityCode string `form:"id" binding:"required"`
}
