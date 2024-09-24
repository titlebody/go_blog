package article_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/res"
)

func (ArticleApi) ArticleRemoveView(c *gin.Context) {
	var cr model.RemoveRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	// 查询文章列表
	var articleModel []model.ArticleModel
	count := global.DB.Find(&articleModel, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("文章不存在", c)
		return
	}

	// 获取所有文章的 ID 列表
	articleIDs := make([]uint, len(articleModel))
	for i, article := range articleModel {
		articleIDs[i] = article.ID
	}

	// 删除 article_tag_models 中的相关记录
	err := global.DB.Model(&articleModel).Association("TagModels").Clear()
	if err != nil {
		global.Log.Error(err)
		return
	}

	// 删除文章
	err = global.DB.Delete(&articleModel).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除文章失败！", c)
		return
	}

	res.OkWithMessage(fmt.Sprintf("删除成功%d个文章", count), c)
}
