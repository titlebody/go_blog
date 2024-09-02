package core

import (
	"go_blog/config"
	"go_blog/global"
	"gopkg.in/yaml.v2"
	"io/fs"
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

func SetYaml() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		global.Log.Error(err)
		return err
	}
	err = ioutil.WriteFile("./config.yaml", byteData, fs.ModePerm)
	if err != nil {
		global.Log.Error(err)
		return err
	}
	global.Log.Info("配置文件修改成功！")
	return nil
}
