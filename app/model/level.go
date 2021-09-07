/**
 *
 * @author 摆渡人
 * @since 2021/8/20
 * @File : level
 */
package model

import (
	"easygoadmin/library/db"
	"time"
)

//映射数据表
func TableName() string {
	return "sys_level"
}

// 对象结构体
type Level struct {
	Id         int       `xorm:"not null pk autoincr comment('职级ID') INT(11)"    json:"id"` // 主键ID
	Name       string    `xorm:"name"        json:"name"`                                   // 职级名称
	Status     int       `xorm:"status"      json:"status"`                                 // 状态：1正常 2停用
	Sort       int       `xorm:"sort"        json:"sort"`                                   // 显示顺序
	CreateUser int       `xorm:"create_user" json:"createUser"`                             // 添加人
	CreateTime time.Time `xorm:"create_time" json:"createTime"`                             // 创建时间
	UpdateUser int       `xorm:"update_user" json:"updateUser"`                             // 更新人
	UpdateTime time.Time `xorm:"update_time" json:"updateTime"`                             // 更新时间
	Mark       int       `xorm:"mark"        json:"mark"`                                   // 有效标识
}

// 分页查询
type LevelPageReq struct {
	Name  string `form:"name"`  // 职级名称
	Page  int    `form:"page"`  // 页码
	Limit int    `form:"limit"` // 每页数
}

// 添加职级
type LevelAddReq struct {
	Name   string `form:"name"  binding:"required#职级名称不能为空"`
	Status int    `form:"status"    binding:"required#职级状态不能为空"`
	Sort   int    `form:"sort"  binding:"required#显示顺序不能为空"`
}

// 编辑职级
type LevelUpdateReq struct {
	Id     int    `form:"id" binding:"required"`
	Name   string `form:"name"  binding:"required"`
	Status int    `form:"status"    binding:"required"`
	Sort   int    `form:"sort"  binding:"required"`
}

// 删除职级
type LevelDeleteReq struct {
	Ids string `form:"ids"  binding:"required#请选择要删除的数据记录"`
}

// 设置状态
type LevelStatusReq struct {
	Id     int `form:"id" binding:"required#主键ID不能为空"`
	Status int `form:"status"    binding:"required#职级状态不能为空"`
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *Level) FindOne() (bool, error) {
	return db.Instance().Engine().Table(TableName()).Get(r)
}

// 插入数据
func (r *Level) Insert() (int64, error) {
	return db.Instance().Engine().Table(TableName()).Insert(r)
}

// 更新数据
func (r *Level) Update() (int64, error) {
	return db.Instance().Engine().Table(TableName()).ID(r.Id).Update(r)
}

// 删除
func (r *Level) Delete() (int64, error) {
	return db.Instance().Engine().Table(TableName()).ID(r.Id).Delete(r)
}

////批量删除
//func (r *Level) BatchDelete(ids ...int64) (int64, error) {
//	return 0, nil
//	//return db.Instance().Engine().Table(TableName()).In("id", ids).Delete(r)
//}
