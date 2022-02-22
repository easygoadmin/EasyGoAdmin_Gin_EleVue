// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 EasyGoAdmin深圳研发中心
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
 * 会员-服务类
 * @author 半城风雨
 * @since 2021/11/13
 * @File : member
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

var Member = new(memberService)

type memberService struct{}

func (s *memberService) GetList(req *dto.MemberPageReq) ([]vo.MemberInfoVo, int64, error) {
	// 创建查询实例
	query := utils.XormDb.Where("mark=1")
	// 查询条件
	if req != nil {
		// 用户名
		if req.Username != "" {
			query = query.Where("username=?", req.Username)
		}
		// 性别
		if req.Gender > 0 {
			query = query.Where("gender=?", req.Gender)
		}
	}
	// 排序
	query = query.OrderBy("id desc")
	// 分页
	offset := (req.Page - 1) * req.Limit
	query = query.Limit(req.Limit, offset)
	// 对象转换
	var list []model.Member
	count, err := query.FindAndCount(&list)
	if err != nil {
		return nil, 0, err
	}

	// 数据处理
	var result = make([]vo.MemberInfoVo, 0)
	for _, v := range list {
		item := vo.MemberInfoVo{}
		item.Member = v
		// 头像
		if v.Avatar != "" {
			item.Avatar = utils.GetImageUrl(v.Avatar)
		}
		// 性别
		if v.Gender > 0 {
			item.GenderName = utils.GENDER_LIST[v.Gender]
		}
		// 设备类型
		if v.Device > 0 {
			item.DeviceName = common.MEMBER_DEVICE_LIST[v.Device]
		}
		// 会员来源
		if v.Source > 0 {
			item.SourceName = common.MEMBER_SOURCE_LIST[v.Source]
		}
		// 所属城市
		if v.DistrictCode != "" {
			item.CityName = City.GetCityName(v.DistrictCode, ">>")
		}
		// 省市区
		cityList := make([]string, 0)
		// 省份编号
		cityList = append(cityList, item.ProvinceCode)
		// 城市编号
		cityList = append(cityList, item.CityCode)
		// 县区编号
		cityList = append(cityList, item.DistrictCode)
		item.City = cityList
		// 加入数组
		result = append(result, item)
	}
	return result, count, nil
}

func (s *memberService) Add(req *dto.MemberAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 实例化对象
	var entity model.Member
	entity.Username = req.Username
	entity.MemberLevel = req.MemberLevel
	entity.Realname = req.Realname
	entity.Nickname = req.Nickname
	entity.Gender = req.Gender
	entity.Birthday = req.Birthday
	entity.Address = req.Address
	entity.Intro = req.Intro
	entity.Signature = req.Signature
	entity.Device = req.Device
	entity.Source = req.Source
	entity.Status = req.Status
	entity.CreateUser = userId
	entity.CreateTime = time.Now().Unix()
	entity.Mark = 1

	// 头像处理
	if req.Avatar != "" {
		avatar, err := utils.SaveImage(req.Avatar, "member")
		if err != nil {
			return 0, err
		}
		entity.Avatar = avatar
	}

	// 省市区处理
	if len(req.City) == 3 {
		entity.ProvinceCode = req.City[0]
		entity.CityCode = req.City[1]
		entity.DistrictCode = req.City[2]
	}

	// 密码
	if req.Password != "" {
		password, _ := utils.Md5(req.Password + req.Username)
		entity.Password = password
	}

	// 插入数据
	return entity.Insert()
}

func (s *memberService) Update(req *dto.MemberUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录
	entity := &model.Member{Id: req.Id}
	has, err := entity.Get()
	if err != nil || !has {
		return 0, err
	}
	entity.Username = req.Username
	entity.MemberLevel = req.MemberLevel
	entity.Realname = req.Realname
	entity.Nickname = req.Nickname
	entity.Gender = req.Gender
	entity.Birthday = req.Birthday
	entity.Address = req.Address
	entity.Intro = req.Intro
	entity.Signature = req.Signature
	entity.Device = req.Device
	entity.Source = req.Source
	entity.Status = req.Status
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now().Unix()

	// 头像处理
	if req.Avatar != "" {
		avatar, err := utils.SaveImage(req.Avatar, "member")
		if err != nil {
			return 0, err
		}
		entity.Avatar = avatar
	}

	// 省市区处理
	if len(req.City) == 3 {
		entity.ProvinceCode = req.City[0]
		entity.CityCode = req.City[1]
		entity.DistrictCode = req.City[2]
	}

	// 密码
	if req.Password != "" {
		password, _ := utils.Md5(req.Password + req.Username)
		entity.Password = password
	}

	// 调用更新方法
	return entity.Update()
}

func (s *memberService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &model.Member{Id: gconv.Int(ids)}
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

func (s *memberService) Status(req *dto.MemberStatusReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录是否存在
	info := &model.Member{Id: req.Id}
	has, err := info.Get()
	if err != nil || !has {
		return 0, errors.New("记录不存在")
	}

	// 设置状态
	entity := &model.Member{}
	entity.Id = info.Id
	entity.Status = req.Status
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now().Unix()
	return entity.Update()
}
