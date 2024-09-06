package config

type QiNiu struct {
	IsEnable  bool    `yaml:"is_enable" json:"is_enable"` //是否启用七牛云存储
	AccessKey string  `yaml:"access_key" json:"access_key"`
	SecretKey string  `yaml:"secret_key" json:"secret_key"`
	Bucket    string  `yaml:"bucket" json:"bucket"` // 存储桶名字
	CDN       string  `yaml:"cdn" json:"cdn"`       //访问图片地址的前缀
	Zone      string  `yaml:"zone" json:"zone"`     // 存储的地区
	Size      float64 `yaml:"size" json:"size"`     // 存储大小限制单位mb
	Prefix    string  `yaml:"prefix" json:"prefix"` // 文件前缀
}
