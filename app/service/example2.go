// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ 赋能开发者，助力企业发展 ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 深圳EasyGoAdmin研发中心
// +----------------------------------------------------------------------
// | Licensed LGPL-3.0 EasyGoAdmin并不是自由软件，未经许可禁止去掉相关版权
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: @半城风雨 团队荣誉出品
// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本团队对该软件框架产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架禁止任何单位和个人、组织用于
// | 任何违法、侵害他人合法利益等恶意的行为，如有发现违规、违法的犯罪行为，本团队将无条件配
// | 合公安机关调查取证同时保留一切以法律手段起诉的权利，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

/**
 * 演示二管理-服务类
 * @author 半城风雨
 * @since 2022-05-16
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
		
		// 演示名称
		
		if req.Name != "" {
			query = query.Where("name like ?", req.Name)
		}
		
	
		// 状态：1正常 2停用
		
		if req.Status > 0 {
			query = query.Where("status=?", req.Status)
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




