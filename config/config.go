// Package config 读取配置文件
package config

type Config struct {
	System   System   `yaml:"system"`
	Logger   Logger   `yaml:"logger"`
	Mysql    Mysql    `yaml:"mysql"`
	SiteInfo SiteInfo `yaml:"site_info"`
	Email    Email    `yaml:"email"`
	JWT      JWT      `yaml:"jwt"`
	QiNiu    QiNiu    `yaml:"qi_niu"`
	QQ       QQ       `yaml:"qq"`
	Upload   Upload   `yaml:"upload"`
	Redis    Redis    `yaml:"redis"`
	ES       ES       `yaml:"es"`
}
