package model

// AdvertModel 广告模型
type AdvertModel struct {
	Model
	Title  string `gorm:"size:32" json:"title"` // 标题
	Href   string `json:"href"`                 // 跳转链接地址
	Images string `json:"images"`               // 图片地址
	IsShow bool   `json:"is_show"`              // 是否显示
}
