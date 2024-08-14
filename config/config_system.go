package config

import "fmt"

// System 系统配置
type system struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
}

func (s system) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
