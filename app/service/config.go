/**
 *
 * @author 半城风雨
 * @since 2021/11/13
 * @File : config
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

var Config = new(configService)

type configService struct{}

func (s *configService) GetList() []model.Config {
	// 创建查询对象
	query := utils.XormDb.
		Where("mark=1").
		OrderBy("sort asc")
	// 对象转换
	var list []model.Config
	query.FindAndCount(&list)
	return list
}

func (s *configService) Add(req *dto.ConfigAddReq, userId int) (int64, error) {
	// 实例化对象
	var entity model.Config
	entity.Name = req.Name
	entity.Sort = req.Sort
	entity.CreateUser = userId
	entity.CreateTime = time.Now()
	entity.Mark = 1

	// 插入记录
	return entity.Insert()
}

func (s *configService) Update(req *dto.ConfigUpdateReq, userId int) (int64, error) {
	// 查询记录
	entity := &model.Config{Id: req.Id}
	has, err := entity.Get()
	if err != nil || !has {
		return 0, err
	}

	// 设置对象
	entity.Name = req.Name
	entity.Sort = req.Sort
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()

	// 更新数据
	return entity.Update()
}

func (s *configService) Delete(ids string) (int64, error) {
	// 记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &model.Config{Id: gconv.Int(ids)}
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
