/**
 *
 * @author 半城风雨
 * @since 2021/11/13
 * @File : item_cate
 */
package dto

// 栏目查询条件
type ItemCateQueryReq struct {
	Name string `form:"name"` // 栏目名称
}

// 添加站点
type ItemCateAddReq struct {
	Name    string `form:"name"        binding:"required"` // 栏目名称
	Pid     int    `form:"pid"`                            // 父级ID
	ItemId  int    `form:"item_id"     binding:"required"` // 站点ID
	Pinyin  string `form:"pinyin"      binding:"required"` // 拼音(全)
	Code    string `form:"code"        binding:"required"` // 拼音(简)
	IsCover int    `form:"is_cover"    binding:"required"` // 是否有封面：1是 2否
	Cover   string `form:"cover"`                          // 封面
	Status  int    `form:"status"      binding:"required"` // 状态：1启用 2停用
	Note    string `form:"note"`                           // 备注
	Sort    int    `form:"sort"        binding:"required"` // 排序
}

// 修改
type ItemCateUpdateReq struct {
	Id      int    `form:"id" binding:"required"`
	Name    string `form:"name"        binding:"required"` // 栏目名称
	Pid     int    `form:"pid"`                            // 父级ID
	ItemId  int    `form:"item_id"     binding:"required"` // 站点ID
	Pinyin  string `form:"pinyin"      binding:"required"` // 拼音(全)
	Code    string `form:"code"        binding:"required"` // 拼音(简)
	IsCover int    `form:"is_cover"    binding:"required"` // 是否有封面：1是 2否
	Cover   string `form:"cover"`                          // 封面
	Status  int    `form:"status"      binding:"required"` // 状态：1启用 2停用
	Note    string `form:"note"`                           // 备注
	Sort    int    `form:"sort"        binding:"required"` // 排序
}
