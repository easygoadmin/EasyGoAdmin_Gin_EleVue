/**
 *
 * @author 半城风雨
 * @since 2021/11/13
 * @File : dict_data
 */
package service

import (
	"easygoadmin/app/dto"
	"easygoadmin/app/model"
	"easygoadmin/utils"
	"easygoadmin/utils/gconv"
	"errors"
	"strings"
	"time"
)

var DictData = new(dictDataService)

type dictDataService struct{}

func (s *dictDataService) GetList(req *dto.DictDataPageReq) ([]model.DictData, int64, error) {
	// 创建查询对象
	query := utils.XormDb.Where("dict_id=?", req.DictId).Where("mark=1")
	// 查询条件
	if req != nil {
		// 字典项名称
		if req.Name != "" {
			query = query.Where("name like ?", "%"+req.Name+"%")
		}
	}
	// 排序
	query = query.OrderBy("sort asc")
	// 分页
	offset := (req.Page - 1) * req.Limit
	query = query.Limit(req.Limit, offset)
	// 对象转换
	var list []model.DictData
	count, err := query.FindAndCount(&list)
	return list, count, err
}

func (s *dictDataService) Add(req *dto.DictDataAddReq, userId int) (int64, error) {
	// 实例化对象
	var entity model.DictData
	entity.DictId = req.DictId
	entity.Name = req.Name
	entity.Code = req.Code
	entity.Sort = req.Sort
	entity.Note = req.Note
	entity.CreateUser = userId
	entity.CreateTime = time.Now()
	entity.Mark = 1

	// 插入记录
	return entity.Insert()
}

func (s *dictDataService) Update(req *dto.DictDataUpdateReq, userId int) (int64, error) {
	// 查询记录
	entity := &model.DictData{Id: req.Id}
	has, err := entity.Get()
	if err != nil || !has {
		return 0, err
	}
	// 设置对象
	entity.DictId = req.DictId
	entity.Name = req.Name
	entity.Code = req.Code
	entity.Sort = req.Sort
	entity.Note = req.Note
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()

	// 更新记录
	return entity.Update()
}

func (s *dictDataService) Delete(ids string) (int64, error) {
	// 记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &model.DictData{Id: gconv.Int(ids)}
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
