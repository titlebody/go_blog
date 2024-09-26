package article_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/res"
)

func (ArticleApi) ArticleSearchView(c *gin.Context) {
	// 文章的模糊搜索
	title := c.Query("title")

	// 进行模糊查询
	var articles []model.ArticleModel
	err := global.DB.Where("title LIKE ? OR content LIKE ?", "%"+title+"%", "%"+title+"%").Find(&articles).Error
	if err != nil {
		res.FailWithMessage("文章不存在", c)
		return
	}
	res.OkWithData(articles, c)
}
