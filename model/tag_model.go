package model

type TagModel struct {
	Model
	Title    string         `gorm:"size:16" json:"title"`
	Articles []ArticleModel `gorm:"many2many:article_tag " json:"articles"`
}
