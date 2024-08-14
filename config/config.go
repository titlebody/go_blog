// Package config 读取配置文件
package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type config struct {
	System system `yaml:"system"`
	Logger logger `yaml:"logger"`
	Mysql  mysql  `yaml:"mysql"`
}

var Config *config

func InitConfig() {
	yamlFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		return
	}
	yaml.Unmarshal(yamlFile, &Config)
}
