package tag_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/res"
)

func (TagApi) TagRemoveView(c *gin.Context) {
	var cr model.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var tagList []model.TagModel
	count := global.DB.Find(&tagList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("标签不存在", c)
		return
	}
	// 如果标签下有文章怎么办
	global.DB.Delete(&tagList)
	res.OkWithMessage(fmt.Sprintf("共删除 %d 个标签", count), c)
}
