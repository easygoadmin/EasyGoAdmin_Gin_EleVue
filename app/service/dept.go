/**
 *
 * @author 半城风雨
 * @since 2021/9/13
 * @File : dept
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

var Dept = new(deptService)

type deptService struct{}

func (s *deptService) GetList(req *dto.DeptPageReq) ([]model.Dept, error) {
	// 创建查询实例
	query := utils.XormDb.Where("mark=1")
	// 查询条件
	if req != nil {
		// 部门名称
		if req.Name != "" {
			query = query.Where("name like ?", "%"+req.Name+"%")
		}
	}
	// 排序
	query = query.OrderBy("sort")
	// 查询数据
	var list []model.Dept
	err := query.Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *deptService) Add(req *dto.DeptAddReq, userId int) (int64, error) {
	// 实例化对象
	var entity model.Dept
	entity.Name = req.Name
	entity.Code = req.Code
	entity.Fullname = req.Fullname
	entity.Type = req.Type
	entity.Pid = req.Pid
	entity.Sort = req.Sort
	entity.Note = req.Note
	entity.CreateUser = userId
	entity.CreateTime = time.Now()
	entity.Mark = 1
	// 插入记录
	rows, err := entity.Insert()
	if err != nil || rows == 0 {
		return 0, errors.New("添加失败")
	}
	return rows, nil
}

func (s *deptService) Update(req *dto.DeptUpdateReq, userId int) (int64, error) {
	// 查询记录
	entity := &model.Dept{Id: req.Id}
	has, err := entity.Get()
	if err != nil || !has {
		return 0, err
	}
	// 设置参数
	entity.Name = req.Name
	entity.Code = req.Code
	entity.Fullname = req.Fullname
	entity.Type = req.Type
	entity.Pid = req.Pid
	entity.Sort = req.Sort
	entity.Note = req.Note
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	// 更新记录
	rows, err := entity.Update()
	if err != nil || rows == 0 {
		return 0, err
	}
	return rows, nil
}

func (s *deptService) Delete(ids string) (int64, error) {
	// 记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := model.Dept{Id: gconv.Int(ids)}
		rows, err := entity.Delete()
		if err != nil || rows == 0 {
			return 0, err
		}
		return rows, nil
	} else {
		// 批量删除
		return 0, nil
	}
}
