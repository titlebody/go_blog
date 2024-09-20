package message_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/model"
	"go_blog/model/res"
	"go_blog/service/common"
)

func (MessageApi) MessageListAllView(c *gin.Context) {
	var cr model.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, _ := common.ComList(model.MessageModel{}, common.Option{
		PageInfo: cr,
	})

	res.OkWithList(list, count, c)

}
