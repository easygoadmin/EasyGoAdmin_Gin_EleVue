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
 * 配置数据-服务类
 * @author 半城风雨
 * @since 2021/11/13
 * @File : config_data
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

var ConfigData = new(configDataService)

type configDataService struct{}

func (s *configDataService) GetList(req *dto.ConfigDataPageReq) ([]vo.ConfigDataVo, int64, error) {
	// 创建查询对象
	query := utils.XormDb.Where("config_id=?", req.ConfigId).Where("mark=1")
	// 查询条件
	if req != nil {
		// 字典项名称
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
	var list []model.ConfigData
	count, err := query.FindAndCount(&list)
	if err != nil {
		return nil, 0, err
	}

	// 数据源处理
	var result = make([]vo.ConfigDataVo, 0)
	for _, v := range list {
		typeName, ok := common.CONFIG_DATA_TYPE_LIST[v.Type]
		item := vo.ConfigDataVo{}
		item.ConfigData = v
		if ok {
			item.TypeName = typeName
		}
		result = append(result, item)
	}
	return result, count, nil
}

func (s *configDataService) Add(req *dto.ConfigDataAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 实例化对象
	var entity model.ConfigData
	entity.Title = req.Title
	entity.Code = req.Code
	entity.Value = req.Value
	entity.Options = req.Options
	entity.ConfigId = req.ConfigId
	entity.Type = req.Type
	entity.Sort = req.Sort
	entity.Note = req.Note
	entity.CreateUser = userId
	entity.CreateTime = time.Now().Unix()
	entity.Mark = 1

	// 插入数据
	return entity.Insert()
}

func (s *configDataService) Update(req *dto.ConfigDataUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录
	entity := &model.ConfigData{Id: req.Id}
	has, err := entity.Get()
	if err != nil || !has {
		return 0, err
	}

	// 设置对象
	entity.Title = req.Title
	entity.Code = req.Code
	entity.Value = req.Value
	entity.Options = req.Options
	entity.ConfigId = req.ConfigId
	entity.Type = req.Type
	entity.Sort = req.Sort
	entity.Note = req.Note
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now().Unix()

	// 更新记录
	return entity.Update()
}

func (s *configDataService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &model.ConfigData{Id: gconv.Int(ids)}
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

func (s *configDataService) Status(req *dto.ConfigDataStatusReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录是否存在
	info := &model.ConfigData{Id: req.Id}
	has, err := info.Get()
	if err != nil || !has {
		return 0, errors.New("记录不存在")
	}

	// 设置状态
	entity := &model.ConfigData{}
	entity.Id = info.Id
	entity.Status = req.Status
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now().Unix()
	return entity.Update()
}
