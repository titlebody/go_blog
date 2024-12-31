package config

type SiteInfo struct {
	CreatedAt   string `yaml:"created_at" json:"created_at" `     // 创建时间
	BeiAn       string `yaml:"bei_an" json:"bei_an" `             // 备案号
	Title       string `yaml:"title" json:"title" `               // 网站标题
	QQImage     string `yaml:"qq_image" json:"qq_image" `         // QQ
	Version     string `yaml:"version" json:"version" `           // 版本
	Email       string `yaml:"email" json:"email" `               // 邮箱
	WechatImage string `yaml:"wechat_image" json:"wechat_image" ` // 微信
	Name        string `yaml:"name" json:"name" `                 // 名字
	Job         string `yaml:"job" json:"job" `                   // 职位
	Addr        string `yaml:"addr" json:"addr" `                 // 地址
	Slogan      string `yaml:"slogan" json:"slogan"`              //  slogan
	SloganEn    string `yaml:"slogan_en" json:"slogan_en" `       // slogan
	Web         string `yaml:"web" json:"web" `                   // 官网
	BiliBiliUrl string `yaml:"bilibili_url" json:"bilibili_url" ` // bili
	GithubUrl   string `yaml:"github_url" json:"github_url" `     // github
}
