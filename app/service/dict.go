/**
 *
 * @author 半城风雨
 * @since 2021/11/13
 * @File : dict
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

var Dict = new(dictService)

type dictService struct{}

func (s *dictService) GetList() []model.Dict {
	// 创建查询对象
	query := utils.XormDb.
		Where("mark=1").
		OrderBy("sort asc")
	// 对象转换
	var list []model.Dict
	query.FindAndCount(&list)
	return list
}

func (s *dictService) Add(req *dto.DictAddReq, userId int) (int64, error) {
	// 实例化对象
	var entity model.Dict
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

func (s *dictService) Update(req *dto.DictUpdateReq, userId int) (int64, error) {
	// 查询记录
	entity := &model.Dict{Id: req.Id}
	has, err := entity.Get()
	if err != nil || !has {
		return 0, err
	}
	if entity == nil {
		return 0, errors.New("记录不存在")
	}

	// 设置对象
	entity.Name = req.Name
	entity.Code = req.Code
	entity.Sort = req.Sort
	entity.Note = req.Note
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()

	return entity.Update()
}

func (s *dictService) Delete(ids string) (int64, error) {
	// 记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &model.Dict{Id: gconv.Int(ids)}
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
