package advert_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/res"
)

// AdvertUpdateView
// @Tags 广告管理
// @Summary 修改广告
// @Description 修改广告
// @Param data body  AdvertRequest true "广告的一些参数"
// @Param id path int true "广告ID"
// @Router /api/adverts/{id} [put]
// @Accept json
// @Success 200 {object} res.Response{data=string}
func (AdvertApi) AdvertUpdateView(c *gin.Context) {
	var cr AdvertRequest
	id := c.Param("id")
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	// 更新广告
	var advert model.AdvertModel
	err := global.DB.Take(&advert, "id = ?", id).Error
	if err != nil {
		res.FailWithMessage("广告不存在", c)
		return
	}
	// 结构体转map的包
	m1 := structs.Map(&cr)
	err = global.DB.Model(&advert).Updates(m1).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改广告失败", c)
		return
	}
	res.OkWithMessage("修改广告成功", c)

}
