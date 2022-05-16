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
 * 友链-服务类
 * @author 半城风雨
 * @since 2021/11/13
 * @File : link
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

var Link = new(linkService)

type linkService struct{}

func (s *linkService) GetList(req *dto.LinkPageReq) ([]vo.LinkInfoVo, int64, error) {
	// 实例化对象
	query := utils.XormDb.Where("mark=1")
	// 查询条件
	if req != nil {
		// 友链名称
		if req.Name != "" {
			query = query.Where("name like ?", "%"+req.Name+"%")
		}
		// 友链类型
		if req.Type > 0 {
			query = query.Where("type=?", req.Type)
		}
		// 投放平台
		if req.Platform > 0 {
			query = query.Where("platform=?", req.Platform)
		}
	}
	// 排序
	query = query.OrderBy("sort asc")
	// 分页
	offset := (req.Page - 1) * req.Limit
	query = query.Limit(req.Limit, offset)
	// 对象转换
	var list []model.Link
	count, err := query.FindAndCount(&list)
	if err != nil {
		return nil, 0, err
	}

	var result []vo.LinkInfoVo
	for _, v := range list {
		item := vo.LinkInfoVo{}
		item.Link = v
		// 友链图片
		if v.Image != "" {
			item.Image = utils.GetImageUrl(v.Image)
		}
		// 友链类型
		typeName, ok := common.LINK_TYPE_LIST[v.Type]
		if ok {
			item.TypeName = typeName
		}
		// 友链形式
		formName, ok := common.LINK_FORM_LIST[v.Form]
		if ok {
			item.FormName = formName
		}
		// 投放平台
		platformName, ok := common.LINK_PLATFORM_LIST[v.Platform]
		if ok {
			item.PlatformName = platformName
		}
		result = append(result, item)
	}
	// 返回结果
	return result, count, nil
}

func (s *linkService) Add(req *dto.LinkAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 实例化对象
	var entity model.Link
	entity.Name = req.Name
	entity.Type = req.Type
	entity.Url = req.Url
	entity.ItemId = req.ItemId
	entity.CateId = req.CateId
	entity.Platform = req.Platform
	entity.Form = req.Form
	entity.Status = req.Status
	entity.Sort = req.Sort
	entity.Note = req.Note
	entity.CreateUser = userId
	entity.CreateTime = time.Now().Unix()
	entity.Mark = 1

	// 图片处理
	if req.Image != "" {
		image, err := utils.SaveImage(req.Image, "link")
		if err != nil {
			return 0, err
		}
		entity.Image = image
	}

	// 插入数据
	return entity.Insert()
}

func (s *linkService) Update(req *dto.LinkUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录
	entity := &model.Link{Id: req.Id}
	has, err := entity.Get()
	if err != nil || !has {
		return 0, err
	}

	// 保存图片
	image, err := utils.SaveImage(req.Image, "link")
	if err != nil {
		return 0, err
	}
	entity.Image = image

	// 设置对象
	entity.Name = req.Name
	entity.Type = req.Type
	entity.Url = req.Url
	entity.ItemId = req.ItemId
	entity.CateId = req.CateId
	entity.Platform = req.Platform
	entity.Form = req.Form
	entity.Status = req.Status
	entity.Sort = req.Sort
	entity.Note = req.Note
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now().Unix()

	// 图片处理
	if req.Image != "" {
		image, err := utils.SaveImage(req.Image, "link")
		if err != nil {
			return 0, err
		}
		entity.Image = image
	}

	// 更新记录
	return entity.Update()
}

func (s *linkService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &model.Link{Id: gconv.Int(ids)}
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

func (s *linkService) Status(req *dto.LinkStatusReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录是否存在
	info := &model.Link{Id: req.Id}
	has, err := info.Get()
	if err != nil || !has {
		return 0, errors.New("记录不存在")
	}

	// 设置状态
	entity := &model.Link{}
	entity.Id = info.Id
	entity.Status = req.Status
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now().Unix()
	return entity.Update()
}
