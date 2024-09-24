package article_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/res"
)

func (ArticleApi) ArticleUpdateView(c *gin.Context) {
	var cr CreateArticleRequest
	if err := c.ShouldBind(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	id := c.Param("id")
	var articleModel model.ArticleModel
	// 使用id查询文章
	err := global.DB.Take(&articleModel, id).Error
	if err != nil {
		res.FailWithMessage("文章不存在", c)
		return
	}

	var banner model.BannerModel
	if cr.BannerID > 0 {
		count := global.DB.Find(&banner, cr.BannerID).RowsAffected
		if count == 0 {
			res.FailWithMessage("文章封面图片不存在", c)
			return
		}
		articleModel.BannerID = cr.BannerID
		articleModel.BannerPath = banner.Path
	}

	// 更新文章
	err = global.DB.Model(&articleModel).Updates(model.ArticleModel{
		Title:      cr.Title,
		Abstract:   cr.Abstract,
		Content:    cr.Content,
		Category:   cr.Category,
		Source:     cr.Source,
		Tags:       cr.Tags,
		NickName:   cr.NickName,
		BannerID:   cr.BannerID,
		BannerPath: banner.Path,
	}).Error
	if err != nil {
		res.FailWithMessage("修改文章失败", c)
		return
	}
	res.OkWithMessage("修改文章成功", c)

}
