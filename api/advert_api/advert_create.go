package advert_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/res"
)

type AdvertRequest struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题" structs:"title"`        // 广告标题
	Href   string `json:"href" binding:"required,url" msg:"广告链接非法" structs:"href"`     // 广告链接
	Images string `json:"images" binding:"required,url" msg:"图片地址非法" structs:"images"` // 广告图片
	IsShow bool   `json:"is_show"  msg:"请选择是否显示" structs:"is_show"`                    // 是否显示
}

// AdvertCreateView
// @Tags 广告管理
// @Summary 创建广告
// @Description 创建广告
// @Param data body AdvertRequest true "表示多个参数"
// @Router /api/adverts [post]
// @Accept json
// @Success 200 {object} res.Response{}
func (AdvertApi) AdvertCreateView(c *gin.Context) {
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	// 重复判断
	var advert model.AdvertModel
	err = global.DB.Take(&advert, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMessage("广告标题重复", c)
		return
	}
	err = global.DB.Create(&model.AdvertModel{
		Title:  cr.Title,
		Href:   cr.Href,
		Images: cr.Images,
		IsShow: cr.IsShow,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("添加广告失败", c)
		return
	}
	res.OkWithMessage("添加广告成功", c)

}
