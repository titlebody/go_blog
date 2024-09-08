package advert_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/res"
)

type AdvertRequest struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题"`       // 广告标题
	Href   string `json:"href" binding:"required,url" msg:"广告链接非法"`   // 广告链接
	Images string `json:"images" binding:"required,url" msg:"图片地址非法"` // 广告图片
	IsShow bool   `json:"is_show" binding:"required" msg:"请选择是否显示"`   // 是否显示
}

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
