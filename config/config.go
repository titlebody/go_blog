// Package config 读取配置文件
package config

type Config struct {
	System system `yaml:"system"`
	Logger logger `yaml:"logger"`
	Mysql  mysql  `yaml:"mysql"`
}
