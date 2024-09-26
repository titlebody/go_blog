package fade_back_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/model"
	"go_blog/model/res"
	"go_blog/service/common"
)

func (FadeBackApi) FadeBackListView(c *gin.Context) {
	var page model.PageInfo
	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, _ := common.ComList(model.FadeBackModel{}, common.Option{
		PageInfo: page,
	})
	// 获取反馈表里的数据
	res.OkWithList(list, count, c)

}
