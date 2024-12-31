package article_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/res"
)

type UpdateArticleRequest struct {
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

func (ArticleApi) ArticleUpdateView(c *gin.Context) {
	var cr UpdateArticleRequest
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
