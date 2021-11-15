/**
 *
 * @author 半城风雨
 * @since 2021/9/8
 * @File : user
 */
package dao

import (
	"easygoadmin/app/model"
	"fmt"
	"xorm.io/xorm"
)

type UserDao struct {
	engine *xorm.Engine
}

//生成一个*UserDao对象
func NewUserDao(engine *xorm.Engine) *UserDao {
	return &UserDao{
		engine: engine,
	}
}

//插入数据
func (userDao *UserDao) InsertUser(users *model.User) (int64, error) {
	return userDao.engine.Insert(users)
}

func (userDao *UserDao) GetUserList(page, size int) []model.User {
	var users []model.User
	// 分页设置
	offset := (page - 1) * size
	err := userDao.engine.Limit(size, offset).Find(&users)
	fmt.Println(err)
	return users
}

func (userDao *UserDao) FindtUsers(page, size int) []model.User {
	offset := (page - 1) * size
	users := make([]model.User, 0)
	err := userDao.engine.Limit(size, offset).Find(&users)
	if err != nil {
		return nil
	} else {
		return users
	}
}
