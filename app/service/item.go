/**
 *
 * @author 半城风雨
 * @since 2021/11/13
 * @File : item
 */
package service

import (
	"easygoadmin/app/dto"
	"easygoadmin/app/model"
	"easygoadmin/app/vo"
	"easygoadmin/utils"
	"easygoadmin/utils/common"
	"easygoadmin/utils/gconv"
	"errors"
	"strings"
	"time"
)

var Item = new(itemService)

type itemService struct{}

func (s *itemService) GetList(req *dto.ItemPageReq) ([]vo.ItemInfoVo, int64, error) {
	// 创建查询实例
	query := utils.XormDb.Where("mark=1")
	// 查询条件
	if req != nil {
		// 站点名称
		if req.Name != "" {
			query = query.Where("name like ?", "%"+req.Name+"%")
		}
		// 站点类型
		if req.Type > 0 {
			query = query.Where("type=?", req.Type)
		}
	}
	// 排序
	query = query.OrderBy("sort asc")
	// 分页
	offset := (req.Page - 1) * req.Limit
	query = query.Limit(req.Limit, offset)
	// 对象转换
	var list []model.Item
	count, err := query.FindAndCount(&list)
	if err != nil {
		return nil, 0, err
	}

	// 数据处理
	var result = make([]vo.ItemInfoVo, 0)
	for _, v := range list {
		item := vo.ItemInfoVo{}
		item.Item = v
		// 站点类型
		typeName, ok := common.ITEM_TYPE_LIST[v.Type]
		if ok {
			item.TypeName = typeName
		}
		// 站点图片
		if v.Image != "" {
			item.Image = utils.GetImageUrl(v.Image)
		}
		result = append(result, item)
	}
	return result, count, nil
}

func (s *itemService) Add(req *dto.ItemAddReq, userId int) (int64, error) {
	// 实例化对象
	var entity model.Item
	entity.Name = req.Name
	entity.Type = req.Type
	entity.Url = req.Url
	entity.Status = req.Status
	entity.Note = req.Note
	entity.Sort = req.Status
	entity.CreateUser = userId
	entity.CreateTime = time.Now()
	entity.Mark = 1

	// 图片处理
	image, err := utils.SaveImage(req.Image, "item")
	if err != nil {
		return 0, err
	}
	entity.Image = image

	// 插入数据
	return entity.Insert()
}

func (s *itemService) Update(req *dto.ItemUpdateReq, userId int) (int64, error) {
	// 查询记录
	entity := &model.Item{Id: req.Id}
	has, err := entity.Get()
	if err != nil || !has {
		return 0, err
	}

	// 设置对象
	entity.Name = req.Name
	entity.Type = req.Type
	entity.Url = req.Url
	entity.Status = req.Status
	entity.Note = req.Note
	entity.Sort = req.Status

	// 图片处理
	image, err := utils.SaveImage(req.Image, "item")
	if err != nil {
		return 0, err
	}
	entity.Image = image
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()

	// 更新记录
	return entity.Update()
}

func (s *itemService) Delete(ids string) (int64, error) {
	// 记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &model.Item{Id: gconv.Int(ids)}
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

func (s *itemService) Status(req *dto.ItemStatusReq, userId int) (int64, error) {
	// 查询记录是否存在
	info := &model.Item{Id: req.Id}
	has, err := info.Get()
	if err != nil || !has {
		return 0, errors.New("记录不存在")
	}

	// 设置状态
	entity := &model.Item{}
	entity.Id = info.Id
	entity.Status = req.Status
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	return entity.Update()
}
