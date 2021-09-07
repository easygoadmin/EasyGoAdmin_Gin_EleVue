package db

import (
	"easygoadmin/library/cfg"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"strings"
	"sync"
	"xorm.io/core"
)

type dbEngine struct {
	master *xorm.Engine //主数据库
	slave  *xorm.Engine //从数据库
}

var (
	instance *dbEngine
	once     sync.Once
)

//初始化数据操作 driver为数据库类型
func Instance(driver ...string) *dbEngine {
	once.Do(func() {

		driverName := "mysql"

		if len(driver) > 0 {
			driverName = driver[0]
		}

		var db dbEngine
		config := cfg.Instance()
		//没有配置从数据库
		if len(config.Database.Slave) == 0 {
			engine, err := xorm.NewEngine(driverName, config.Database.Master)
			if err != nil {
				fmt.Printf("数据库连接错误:%v", err.Error())
				return
			}
			err = engine.Ping()
			if err != nil {
				fmt.Printf("数据库连接错误:%v", err.Error())
				return
			}
			if cfg.Instance().Database.Debug {
				engine.ShowSQL(cfg.Instance().Database.Debug)
				engine.Logger().SetLevel(core.LOG_DEBUG)
			}
			db.master = engine
			instance = &db
		} else {
			master, err := xorm.NewEngine(driverName, config.Database.Master)
			if err != nil {
				return
			}
			if cfg.Instance().Database.Debug {
				master.ShowSQL(cfg.Instance().Database.Debug)
				master.Logger().SetLevel(core.LOG_DEBUG)
			}

			slave, err := xorm.NewEngine(driverName, config.Database.Slave)
			if err != nil {
				return
			}
			slaves := []*xorm.Engine{slave}
			group, err := xorm.NewEngineGroup(master, slaves)
			if cfg.Instance().Database.Debug {
				group.ShowSQL(cfg.Instance().Database.Debug)
				group.Logger().SetLevel(core.LOG_DEBUG)
			}
			db.master = group.Master()
			db.slave = group.Slave()
			instance = &db
		}
	})
	return instance
}

//获取操作实例 如果传入slave 并且成功配置了slave 返回slave orm引擎 否则返回master orm引擎
func (db *dbEngine) Engine(dbType ...string) *xorm.Engine {
	if dbType != nil && len(dbType) > 0 {
		if strings.EqualFold(dbType[0], "slave") {
			if db.slave != nil {
				return db.slave
			}
		}
	}
	return db.master
}
