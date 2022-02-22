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
 * 数据库连接工具
 * @author 半城风雨
 * @since 2021/9/8
 * @File : db
 */
package utils

import (
	"easygoadmin/library/cfg"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
	"xorm.io/core"
)

var XormDb *xorm.Engine

func init() {
	fmt.Println("初始化并连接数据库")

	// 获取配置实例
	config := cfg.Instance()
	var err error
	XormDb, err = xorm.NewEngine("mysql", config.Database.Master)
	if err != nil {
		fmt.Printf("数据库连接错误:%v", err.Error())
		return
	}

	// 通过engine.Ping()来进行数据库的连接测试是否可以连接到数据库。
	err = XormDb.Ping()
	if err == nil {
		fmt.Println("数据库连接成功")
		//关闭连接
		//defer XormDb.Close()
	} else {
		fmt.Printf("数据库连接错误:%v", err.Error())
		return
	}

	XormDb.DatabaseTZ = time.Local // 必须
	XormDb.TZLocation = time.Local // 必须
	// 设置连接池的空闲数大小
	XormDb.SetMaxIdleConns(10)
	// 设置最大打开连接数
	XormDb.SetMaxOpenConns(30)

	// 结构体与数据表的映射
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "sys_")
	XormDb.SetTableMapper(tbMapper)

	// 开启调试模式和打印日志,会在控制台打印执行的sql
	if cfg.Instance().Database.Debug {
		XormDb.ShowSQL(cfg.Instance().Database.Debug)
		XormDb.Logger().SetLevel(core.LOG_DEBUG)
	}
}
