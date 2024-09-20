package tag_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/res"
)

func (TagApi) TagUpdateView(c *gin.Context) {
	var cr TagRequest
	id := c.Param("id")
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	// 更新广告
	var tag model.TagModel
	err := global.DB.Take(&tag, "id = ?", id).Error
	if err != nil {
		res.FailWithMessage("标签不存在", c)
		return
	}
	// 结构体转map的包
	m1 := structs.Map(&cr)
	err = global.DB.Model(&tag).Updates(m1).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改标签失败", c)
		return
	}
	res.OkWithMessage("修改标签成功", c)

}
