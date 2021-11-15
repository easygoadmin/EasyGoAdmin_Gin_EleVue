/**
 *
 * @author 半城风雨
 * @since 2021/11/15
 * @File : ad
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

var Ad = new(adService)

type adService struct{}

func (s *adService) GetList(req *dto.AdPageReq) ([]vo.AdInfoVo, int64, error) {
	// 创建查询实例
	query := utils.XormDb.Where("mark=1")
	// 查询条件
	if req != nil {
		// 广告标题
		if req.Title != "" {
			query = query.Where("title like ?", "%"+req.Title+"%")
		}
	}
	// 排序
	query = query.OrderBy("sort asc")
	// 分页
	offset := (req.Page - 1) * req.Limit
	query = query.Limit(req.Limit, offset)
	// 对象转换
	var list []model.Ad
	count, err := query.FindAndCount(&list)
	if err != nil {
		return nil, 0, err
	}

	// 数据处理
	result := make([]vo.AdInfoVo, 0)
	for _, v := range list {
		item := vo.AdInfoVo{}
		item.Ad = v
		// 广告封面
		if v.Cover != "" {
			item.Cover = utils.GetImageUrl(v.Cover)
		}
		// 广告类型
		if v.Type > 0 {
			item.TypeName = common.AD_TYPE_LIST[v.Type]
		}
		// 所属广告位
		if v.AdSortId > 0 {
			adSortInfo := &model.AdSort{}
			has, err := utils.XormDb.Id(v.AdSortId).Get(&adSortInfo)
			if err == nil && has {
				item.AdSortDesc = adSortInfo.Description + " >> " + gconv.String(adSortInfo.LocId)
			}
		}
		result = append(result, item)
	}

	// 返回结果
	return result, count, err
}

func (s *adService) Add(req *dto.AdAddReq, userId int) (int64, error) {
	// 实例化对象
	var entity model.Ad
	entity.Title = req.Title
	entity.AdSortId = req.AdSortId
	entity.Type = req.Type
	entity.Description = req.Description
	entity.Content = req.Content
	entity.Url = req.Url
	entity.Width = req.Width
	entity.Height = req.Height
	entity.StartTime = req.StartTime
	entity.EndTime = req.EndTime
	entity.Status = req.Status
	entity.Sort = req.Sort
	entity.CreateUser = userId
	entity.CreateTime = time.Now()
	entity.Mark = 1

	// 广告封面
	if req.Type == 1 {
		// 图片
		cover, err := utils.SaveImage(req.Cover, "ad")
		if err != nil {
			return 0, err
		}
		entity.Cover = cover
	} else {
		entity.Cover = ""
	}

	// 插入数据
	return entity.Insert()
}

func (s *adService) Update(req *dto.AdUpdateReq, userId int) (int64, error) {
	// 查询记录
	entity := &model.Ad{Id: req.Id}
	has, err := entity.Get()
	if err != nil || !has {
		return 0, err
	}

	// 设置对象
	entity.Title = req.Title
	entity.AdSortId = req.AdSortId
	entity.Type = req.Type
	entity.Description = req.Description
	entity.Content = req.Content
	entity.Url = req.Url
	entity.Width = req.Width
	entity.Height = req.Height
	entity.StartTime = req.StartTime
	entity.EndTime = req.EndTime
	entity.Status = req.Status
	entity.Sort = req.Sort
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()

	// 广告封面
	if req.Type == 1 {
		// 图片
		cover, err := utils.SaveImage(req.Cover, "ad")
		if err != nil {
			return 0, err
		}
		entity.Cover = cover
	} else {
		entity.Cover = ""
	}

	// 更新信息
	return entity.Update()
}

func (s *adService) Delete(ids string) (int64, error) {
	// 记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &model.Ad{Id: gconv.Int(ids)}
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

func (s *adService) Status(req *dto.AdStatusReq, userId int) (int64, error) {
	// 查询记录是否存在
	info := &model.Ad{Id: req.Id}
	has, err := info.Get()
	if err != nil || !has {
		return 0, errors.New("记录不存在")
	}

	// 设置状态
	entity := &model.Ad{}
	entity.Id = info.Id
	entity.Status = req.Status
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	return entity.Update()
}
