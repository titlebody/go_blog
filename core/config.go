package core

import (
	"go_blog/config"
	"go_blog/global"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// InitConfig 读取配置文件
func InitConfig() {
	Config := &config.Config{}
	yamlFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		return
	}
	yaml.Unmarshal(yamlFile, Config)
	global.Config = Config
}
