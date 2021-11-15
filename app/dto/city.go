/**
 *
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
	Lng      int    `form:"lng"`                            // 城市坐标中心点经度（* 1e6）：如果是中国，此值是 1e7
	Lat      int    `form:"lat"`                            // 城市坐标中心点纬度（* 1e6）
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
	Lng      int    `form:"lng"`                            // 城市坐标中心点经度（* 1e6）：如果是中国，此值是 1e7
	Lat      int    `form:"lat"`                            // 城市坐标中心点纬度（* 1e6）
	Sort     int    `form:"sort"`                           // 排序号
}

// 获取子级城市
type CityChildReq struct {
	CityCode string `form:"id" binding:"required"`
}
