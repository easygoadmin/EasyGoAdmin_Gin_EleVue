package cfg

import (
	"github.com/BurntSushi/toml"
	"os"
	"sync"
)

var (
	instance *config
	once     sync.Once
)

//获取配置文档实例
func Instance() *config {
	once.Do(func() {
		var conf config
		path, _ := os.Getwd()
		filePath := path + "\\config\\config.toml"
		if _, err := toml.DecodeFile(filePath, &conf); err != nil {
			return
		}
		instance = &conf
	})

	return instance
}

type config struct {
	Database    database
	Logger      logger
	EasyGoAdmin easygoadmin
}

type database struct {
	Master string
	Slave  string
	Debug  bool
	Log    string
}

type logger struct {
	Path   string
	Level  uint32
	Stdout bool
}

// 自定义配置
type easygoadmin struct {
	Version string
	Debug   bool
	Image   string
	Uploads string
}
