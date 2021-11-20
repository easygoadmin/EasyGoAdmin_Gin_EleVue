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
 * 演示二管理-服务类
 * @author 半城风雨
 * @since 2021-11-20
 * @File : example2
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

// 中间件管理服务
var Example2 = new(example2Service)

type example2Service struct{}

func (s *example2Service) GetList(req *dto.Example2PageReq) ([]vo.Example2InfoVo, int64, error) {
	// 初始化查询实例
	query := utils.XormDb.Where("mark=1")
	if req != nil {
		// 职级名称查询
		if req.Name != "" {
			query = query.Where("name like ?", "%"+req.Name+"%")
		}
	}
	// 排序
	query = query.Asc("sort")
	// 分页设置
	offset := (req.Page - 1) * req.Limit
	query = query.Limit(req.Limit, offset)
	// 查询列表
	list := make([]model.Example2, 0)
	count, err := query.FindAndCount(&list)

	// 数据处理
	var result []vo.Example2InfoVo
	for _, v := range list {
		item := vo.Example2InfoVo{}
		item.Example2 = v

		result = append(result, item)
	}

	// 返回结果
	return result, count, err
}

func (s *example2Service) Add(req *dto.Example2AddReq, userId int) (int64, error) {
	// 实例化对象
	var entity model.Example2

	entity.Name = req.Name
	entity.Status = req.Status
	entity.Sort = req.Sort
	entity.CreateUser = userId
	entity.CreateTime = time.Now().Unix()
	entity.Mark = 1
	// 插入数据
	return entity.Insert()
}

func (s *example2Service) Update(req *dto.Example2UpdateReq, userId int) (int64, error) {
	// 查询记录
	entity := &model.Example2{Id: req.Id}
	has, err := entity.Get()
	if err != nil || !has {
		return 0, errors.New("记录不存在")
	}

	entity.Name = req.Name
	entity.Status = req.Status
	entity.Sort = req.Sort
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now().Unix()
	// 更新记录
	return entity.Update()
}

// 删除
func (s *example2Service) Delete(ids string) (int64, error) {
	// 记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &model.Example2{Id: gconv.Int(ids)}
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

func (s *example2Service) Status(req *dto.Example2StatusReq, userId int) (int64, error) {
	// 查询记录是否存在
	info := &model.Example2{Id: req.Id}
	has, err := info.Get()
	if err != nil || !has {
		return 0, errors.New("记录不存在")
	}

	// 设置状态
	entity := &model.Example2{}
	entity.Id = info.Id
	entity.Status = req.Status
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now().Unix()
	return entity.Update()
}
