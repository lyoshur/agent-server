package config

import (
	"github.com/lyoshur/gutils"
	"strings"
)

// 加载配置文件
func Load() Config {
	// 设置默认配置文件所在位置
	configPath := "./config.xml"
	// 查看是否存在命令行参数指定配置文件
	path, ok := gutils.LoadConfigFromCMD()["config"]
	if ok {
		configPath = path
	}
	// 初始化配置
	conf := Config{}
	// 判断配置文件类型
	if strings.HasPrefix(configPath, "http") {
		gutils.LoadXmlConfigFromHTTP(configPath, &conf)
	} else {
		gutils.LoadXmlConfigFromFile(configPath, &conf)
	}
	return conf
}
