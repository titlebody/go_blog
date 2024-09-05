package images_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/model"
	"go_blog/model/res"
	"go_blog/service/common"
)

func (ImagesApi) ImageListView(c *gin.Context) {
	var cr model.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, err := common.ComList(model.BannerModel{}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	res.OkWithList(list, count, c)
	return
}
