package images_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/res"
)

type ImageUpdateRequest struct {
	ID   uint   `json:"id" binding:"required" msg:"请输入文件ID"`
	Name string `json:"name" binding:"required" msg:"请输入文件名称"`
}

// ImageUpdateView 图片名称修改
func (ImagesApi) ImageUpdateView(c *gin.Context) {
	var cr ImageUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var imageModel model.BannerModel
	err = global.DB.Take(&imageModel, cr.ID).Error
	if err != nil {
		res.FailWithMessage("文件不存在", c)
		return
	}
	err = global.DB.Model(&imageModel).Update("name", cr.Name).Error
	if err != nil {
		res.FailWithMessage("文件修改失败", c)
		return
	}
	res.OkWithMessage("图片名称修改成功", c)
	return
}
