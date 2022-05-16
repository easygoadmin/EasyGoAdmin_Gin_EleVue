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
 * 栏目-服务类
 * @author 半城风雨
 * @since 2021/11/13
 * @File : item_cate
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

var ItemCate = new(itemCateService)

type itemCateService struct{}

func (s *itemCateService) GetList(req *dto.ItemCateQueryReq) []vo.ItemCateInfoVo {
	// 创建查询对象
	query := utils.XormDb.Where("mark=1")
	// 查询条件
	if req != nil {
		// 栏目名称
		if req.Name != "" {
			query = query.Where("name like ?", "%"+req.Name+"%")
		}
	}
	// 排序
	query = query.OrderBy("sort asc")
	// 对象转换
	var list []model.ItemCate
	query.Find(&list)

	// 数据处理
	var result []vo.ItemCateInfoVo
	for _, v := range list {
		item := vo.ItemCateInfoVo{}
		item.ItemCate = v
		// 站点封面
		if v.IsCover == 1 && v.Cover != "" {
			item.Cover = utils.GetImageUrl(v.Cover)
		}
		// 获取栏目
		if v.ItemId > 0 {
			var itemInfo model.Item
			utils.XormDb.Id(item.Id).Get(&itemInfo)
			item.ItemName = itemInfo.Name
		}
		// 加入数组
		result = append(result, item)
	}
	return result
}

func (s *itemCateService) Add(req *dto.ItemCateAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 实例化对象
	var entity model.ItemCate
	entity.Name = req.Name
	entity.Pid = req.Pid
	entity.ItemId = req.ItemId
	entity.Pinyin = req.Pinyin
	entity.Code = req.Code
	entity.Status = req.Status
	entity.Note = req.Note
	entity.Sort = req.Sort

	// 封面
	entity.IsCover = req.IsCover
	if req.IsCover == 1 {
		// 有封面
		cover, err := utils.SaveImage(req.Cover, "item_cate")
		if err != nil {
			return 0, err
		}
		entity.Cover = cover
	} else {
		// 没封面
		entity.Cover = ""
	}
	entity.CreateUser = userId
	entity.CreateTime = time.Now().Unix()
	entity.Mark = 1

	// 插入数据
	return entity.Insert()
}

func (s *itemCateService) Update(req *dto.ItemCateUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录
	entity := &model.ItemCate{Id: req.Id}
	has, err := entity.Get()
	if err != nil || !has {
		return 0, err
	}

	// 设置对象
	entity.Name = req.Name
	entity.Pid = req.Pid
	entity.ItemId = req.ItemId
	entity.Pinyin = req.Pinyin
	entity.Code = req.Code
	entity.Status = req.Status
	entity.Note = req.Note
	entity.Sort = req.Sort

	// 封面
	entity.IsCover = req.IsCover
	if req.IsCover == 1 {
		// 有封面
		cover, err := utils.SaveImage(req.Cover, "item_cate")
		if err != nil {
			return 0, err
		}
		entity.Cover = cover
	} else {
		// 没封面
		entity.Cover = ""
	}
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now().Unix()

	// 更新记录
	return entity.Update()
}

func (s *itemCateService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &model.ItemCate{Id: gconv.Int(ids)}
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

func (s *itemCateService) GetCateName(cateId int, delimiter string) string {
	// 声明数组
	list := make([]string, 0)
	for {
		if cateId <= 0 {
			// 退出
			break
		}
		// 业务处理
		var info model.ItemCate
		has, err := utils.XormDb.Id(cateId).Get(&info)
		if err != nil || !has {
			break
		}
		// 上级栏目ID
		cateId = info.Pid
		// 加入数组
		list = append(list, info.Name)
	}
	// 结果数据处理
	if len(list) > 0 {
		// 数组反转
		utils.Reverse(&list)
		// 拼接字符串
		return strings.Join(list, delimiter)
	}
	return ""
}
