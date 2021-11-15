/**
 *
 * @author 半城风雨
 * @since 2021/11/13
 * @File : item_cate
 */
package service

import (
	"easygoadmin/app/dto"
	"easygoadmin/app/model"
	"easygoadmin/app/vo"
	"easygoadmin/utils"
	"easygoadmin/utils/gconv"
	"errors"
	"strings"
	"time"
)

var ItemCate = new(itemCateService)

type itemCateService struct{}

func (s *itemCateService) GetList(req *dto.ItemCateQueryReq) []vo.ItemCateInfoVo {
	// 创建查询对象
	query := utils.XormDb.Where("mark=1")
	// 查询条件
	if req != nil {
		// 栏目名称
		if req.Name != "" {
			query = query.Where("name like ?", "%"+req.Name+"%")
		}
	}
	// 排序
	query = query.OrderBy("sort asc")
	// 对象转换
	var list []model.ItemCate
	query.Find(&list)

	// 数据处理
	var result []vo.ItemCateInfoVo
	for _, v := range list {
		item := vo.ItemCateInfoVo{}
		item.ItemCate = v
		// 站点封面
		if v.IsCover == 1 && v.Cover != "" {
			item.Cover = utils.GetImageUrl(v.Cover)
		}
		// 获取栏目
		if v.ItemId > 0 {
			var itemInfo model.Item
			utils.XormDb.Id(item.Id).Get(&itemInfo)
			item.ItemName = itemInfo.Name
		}
		// 加入数组
		result = append(result, item)
	}
	return result
}

func (s *itemCateService) Add(req *dto.ItemCateAddReq, userId int) (int64, error) {
	// 实例化对象
	var entity model.ItemCate
	entity.Name = req.Name
	entity.Pid = req.Pid
	entity.ItemId = req.ItemId
	entity.Pinyin = req.Pinyin
	entity.Code = req.Code
	entity.Status = req.Status
	entity.Note = req.Note
	entity.Sort = req.Sort

	// 封面
	entity.IsCover = req.IsCover
	if req.IsCover == 1 {
		// 有封面
		cover, err := utils.SaveImage(req.Cover, "item_cate")
		if err != nil {
			return 0, err
		}
		entity.Cover = cover
	} else {
		// 没封面
		entity.Cover = ""
	}
	entity.CreateUser = userId
	entity.CreateTime = time.Now()
	entity.Mark = 1

	// 插入数据
	return entity.Insert()
}

func (s *itemCateService) Update(req *dto.ItemCateUpdateReq, userId int) (int64, error) {
	// 查询记录
	entity := &model.ItemCate{Id: req.Id}
	has, err := entity.Get()
	if err != nil || !has {
		return 0, err
	}

	// 设置对象
	entity.Name = req.Name
	entity.Pid = req.Pid
	entity.ItemId = req.ItemId
	entity.Pinyin = req.Pinyin
	entity.Code = req.Code
	entity.Status = req.Status
	entity.Note = req.Note
	entity.Sort = req.Sort

	// 封面
	entity.IsCover = req.IsCover
	if req.IsCover == 1 {
		// 有封面
		cover, err := utils.SaveImage(req.Cover, "item_cate")
		if err != nil {
			return 0, err
		}
		entity.Cover = cover
	} else {
		// 没封面
		entity.Cover = ""
	}
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()

	// 更新记录
	return entity.Update()
}

func (s *itemCateService) Delete(ids string) (int64, error) {
	// 记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &model.ItemCate{Id: gconv.Int(ids)}
		rows, err := entity.Delete()
		if err != nil || rows == 0 {
			return 0, errors.New("删除失败")
		}
		return rows, nil
	} else {
		// 批量删除
		return 0, nil
	}
}

func (s *itemCateService) GetCateName(cateId int, delimiter string) string {
	// 声明数组
	list := make([]string, 0)
	for {
		if cateId <= 0 {
			// 退出
			break
		}
		// 业务处理
		var info model.ItemCate
		has, err := utils.XormDb.Id(cateId).Get(&info)
		if err != nil || !has {
			break
		}
		// 上级栏目ID
		cateId = info.Pid
		// 加入数组
		list = append(list, info.Name)
	}
	// 结果数据处理
	if len(list) > 0 {
		// 数组反转
		utils.Reverse(&list)
		// 拼接字符串
		return strings.Join(list, delimiter)
	}
	return ""
}
