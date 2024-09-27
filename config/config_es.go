package config

type ES struct {
	Host string `yaml:"host" json:"host"`
	Port int    `yaml:"port" json:"port"`
	User string `yaml:"user" json:"user"`
	Pass string `yaml:"password" json:"password"`
}
