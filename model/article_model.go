package model

import "go_blog/model/c_type"

// ArticleModel 文章模型
type ArticleModel struct {
	MODEL
	Title         string         `gorm:"size:32" json:"title"`                           //文章标题
	Abstract      string         `json:"abstract"`                                       //文章简介
	Content       string         `json:"content"`                                        //文章内容
	LookCount     int            `json:"look_count"`                                     //阅读次数
	CommentCount  int            `json:"comment_count"`                                  //评论次数
	DiggCount     int            `json:"digg_count"`                                     //点赞次数
	CollectsCount int            `json:"collects_count"`                                 //收藏次数
	TagModels     []TagModel     `gorm:"many2many:article_tag_models" json:"tag_models"` //文章标签
	CommentModels []CommentModel `gorm:"foreignKey:ArticleID" json:"-"`                  //文章评论
	UserModel     UserModel      `gorm:"foreignKey:UserID" json:"-"`                     //文章作者
	UserID        uint           `json:"user_id"`                                        //文章作者ID
	Category      string         `gorm:"size:32" json:"category"`                        //文章分类
	Source        string         `json:"source"`                                         //文章来源
	Link          string         `json:"link"`                                           //文章链接
	Banner        BannerModel    `gorm:"foreignKey:BannerID" json:"-"`                   //文章封面
	BannerID      uint           `json:"banner_id"`                                      //文章封面ID
	NickName      string         `json:"nick_name"`                                      //文章作者昵称
	BannerPath    string         `json:"banner_path"`                                    //文章封面路径
	Tags          c_type.Array   `gorm:"type:string;size:64" json:"tags"`                //文章标签
}
