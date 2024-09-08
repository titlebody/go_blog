package advert_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/model"
	"go_blog/model/res"
	"go_blog/service/common"
	"strings"
)

// AdvertListView
// @Tags 广告管理
// @Summary 广告列表
// @Description 广告列表
// @Param data query model.PageInfo false "查询参数"
// @Router /api/adverts [get]
// @Accept json
// @Success 200 {object} res.Response{data=res.ListResponse[model.AdvertModel]}
func (AdvertApi) AdvertListView(c *gin.Context) {
	var cr model.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	// 判断referer 是否包含admin 如果是，则返回所有广告，不是返回is_show=true
	Referer := c.GetHeader("Referer")
	isShow := true

	if strings.Contains(Referer, "admin") {
		// admin来的
		isShow = false
	}

	list, count, _ := common.ComList(model.AdvertModel{
		IsShow: isShow,
	}, common.Option{
		PageInfo: cr,
	})
	res.OkWithList(list, count, c)

}
