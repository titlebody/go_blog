package model

// TagModel 标签表
type TagModel struct {
	MODEL
	Title        string         `gorm:"size:16" json:"title"`
	Articles     []ArticleModel `gorm:"many2many:article_tag_models" json:"-"`
	ArticleCount int64          `gorm:"-" json:"article_count"`
}
