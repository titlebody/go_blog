package tag_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/res"
)

type TagRequest struct {
	Title string `json:"title" binding:"required" msg:"请输入标题" structs:"title"` // 标签标题
}

func (TagApi) TagCreateView(c *gin.Context) {
	var cr TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	// 重复判断
	var tag model.TagModel
	err = global.DB.Take(&tag, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMessage("该标题重复", c)
		return
	}
	err = global.DB.Create(&model.TagModel{
		Title: cr.Title,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("添加标题失败", c)
		return
	}
	res.OkWithMessage("添加标题成功", c)

}
