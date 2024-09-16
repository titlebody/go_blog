package images_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/res"
)

type ImagesResponse struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
	Name string `json:"name"`
}

// ImageNameListView
// @Tags 图片管理
// @Summary 图片名称列表
// @Description 图片名称列表
// @Router /api/images_name [get]
// @Accept json
// @Success 200 {object} res.Response{ data=[]ImagesResponse}
func (ImagesApi) ImageNameListView(c *gin.Context) {
	var ImageList []ImagesResponse
	global.DB.Model(&model.BannerModel{}).Select("id", "path", "name").Scan(&ImageList)
	res.OkWithData(ImageList, c)
}
