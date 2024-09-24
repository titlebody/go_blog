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
	global.DB.Find(&article, id)
	res.OkWithData(article, c)
}
