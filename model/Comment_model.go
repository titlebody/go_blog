package model

type CommentModel struct {
	Model

	SubComments        []*CommentModel `gorm:"foreignkey:ParentCommentID" json:"sub_comments"`         // 子评论
	ParentCommentModel *CommentModel   `gorm:"foreignkey:ParentCommentID" json:"parent_comment_model"` // 父评论
	ParentCommentId    *uint           `json:"parent_comment_id"`                                      // 父评论ID
	Content            string          `gorm:"size:256" json:"content"`                                // 内容
	DiggCount          int             `gorm:"size:8;default:0" json:"digg_count"`                     // 点赞数
	CommentCount       int             `gorm:"size:8;default:0" json:"comment_count"`                  // 评论数
	Article            ArticleModel    `gorm:"foreignKey:ArticleID" json:"article"`                    // 文章
	ArticleId          uint            `json:"article_id"`                                             // 文章ID
	User               UserModel       `json:"user"`                                                   // 用户
	UserID             uint            `gorm:"size:10" json:"user_id"`                                 // 用户ID
}
