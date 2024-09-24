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

	// 检查标签下是否有文章
	for _, tag := range tagList {
		if global.DB.Preload("Articles").Where("id = ?", tag.ID).First(&tag).RowsAffected > 0 {
			if len(tag.Articles) > 0 {
				res.FailWithMessage("标签下有文章，禁止删除", c)
				return
			}
		}
	}

	global.DB.Delete(&tagList)
	res.OkWithMessage(fmt.Sprintf("共删除 %d 个标签", count), c)
}
