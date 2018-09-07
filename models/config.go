package models

import (
	"log"
	"path/filepath"

	"lxxccc.lib/LxxcccToolkits_golang/config/toml"
	"lxxccc.lib/LxxcccToolkits_golang/file"
)

type ConfigInfo struct {
	RealServer string
}

var Config = ConfigInfo{}

func init() {
	filePath := filepath.Join(file.GetRootPath(), "conf", "config.toml")
	err := toml.ReadTomlConfig(filePath, &Config)
	if nil != err {
		log.Println("缺少配置文件")
	}
}
