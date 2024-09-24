package article_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/res"
	"go_blog/utils/jwts"
)

type CreateArticleRequest struct {
	Title    string   `json:"title" binding:"required" msg:"请输入标题"`
	Abstract string   `json:"abstract" binding:"required" msg:"请输入简介"`
	Content  string   `json:"content" binding:"required" msg:"请输入内容"`
	Category string   `json:"category" binding:"required"  msg:"请输入分类"`
	Source   string   `json:"source"`
	Link     string   `json:"link"`
	NickName string   `json:"nick_name"`
	Tags     []string `json:"tags" binding:"required"  msg:"请输入标签"`
	UserID   uint     `json:"user_id"`
	BannerID uint     `json:"banner_id" msg:"请选择封面" binding:"required"`
}

func (ArticleApi) ArticleCreateView(c *gin.Context) {
	// 获取当前用户数据
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	// 获取请求参数
	var cr CreateArticleRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	// 标题不能重复
	count := global.DB.Find(&model.ArticleModel{}, "title = ?", cr.Title).RowsAffected
	if count > 0 {
		res.OkWithMessage("标题重复", c)
		return
	}

	// 获取 BannerModel 信息
	var banner model.BannerModel
	if err := global.DB.Where("id = ?", cr.BannerID).First(&banner).Error; err != nil {
		return
	}

	article := model.ArticleModel{
		Title:         cr.Title,
		Abstract:      cr.Abstract,
		Content:       cr.Content,
		LookCount:     0,
		CommentCount:  0,
		DiggCount:     0,
		CollectsCount: 0,
		Category:      cr.Category,
		Source:        cr.Source,
		Link:          cr.Link,
		BannerPath:    banner.Path,
		NickName:      claims.NickName,
		UserID:        claims.UserID,
		BannerID:      banner.ID,
		Tags:          cr.Tags,
	}
	global.DB.Create(&article)
	// 获取或创建标签
	var tagModels []model.TagModel
	for _, tagName := range cr.Tags {
		var tag model.TagModel
		global.DB.Where("title = ?", tagName).First(&tag)
		if tag.ID == 0 {
			global.DB.Create(&model.TagModel{Title: tagName})
			global.DB.Where("title = ?", tagName).First(&tag)
		}
		tagModels = append(tagModels, tag)
	}
	// 关联文章和标签
	err := global.DB.Model(&article).Association("TagModels").Append(tagModels)
	if err != nil {
		res.FailWithMessage("关联标签失败", c)
		return
	}
	res.OkWithData(article, c)

}
