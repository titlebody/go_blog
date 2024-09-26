package article_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/res"
)

func (ArticleApi) ArticleView(c *gin.Context) {
	id := c.Param("id")
	var article model.ArticleModel
	err := global.DB.Find(&article, id).Error
	if err != nil {
		res.FailWithMessage("查询失败", c)
		return
	}
	// 每次调用该接口 本文章的浏览量+1
	global.DB.Model(&article).UpdateColumn("look_count", article.LookCount+1)
	res.OkWithData(article, c)

}
