package fade_back_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/res"
)

// UserFadeView 用户申请
//Email        string `gorm:"size:64" json:"email"`
//Content      string `gorm:"size:128" json:"content"`

type FadeBack struct {
	Email   string `json:"email" binding:"required" msg:"请输入邮箱"`
	Content string `json:"content" binding:"required" msg:"请输入内容"`
}

func (FadeBackApi) FadeBackView(c *gin.Context) {
	var cr FadeBack
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	// 数据添加到数据库
	global.DB.Create(&model.FadeBackModel{
		Email:   cr.Email,
		Content: cr.Content,
		IsApply: false,
	})
	res.OkWithMessage("提交成功", c)
}
