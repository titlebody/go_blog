package fade_back_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/res"
)

func (FadeBackApi) RemoveFadeView(c *gin.Context) {
	var cr model.RemoveRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var fadeBackModel []model.FadeBackModel
	count := global.DB.Find(&fadeBackModel, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("不存在", c)
		return
	}
	global.DB.Delete(&fadeBackModel)
	res.OkWithMessage("删除成功", c)

}
