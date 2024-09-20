package message_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/res"
)

type MessageRequest struct {
	SendUserID uint   `json:"send_user_id" binding:"required"`
	RevUserID  uint   `json:"rev_user_id" binding:"required"`
	Content    string `json:"content" binding:"required"` // 内容
}

func (MessageApi) MessageCreateView(c *gin.Context) {
	//当前用户发送消息
	// SendUserID 就是当前发送人的ID

	var cr MessageRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("参数错误", c)
		return
	}
	var sendUser, revUser model.UserModel
	err = global.DB.Take(&sendUser, "id = ?", cr.SendUserID).Error
	if err != nil {
		fmt.Println(2)
		res.FailWithMessage("发送人不存在", c)
		return
	}
	err = global.DB.Take(&revUser, "id = ?", cr.RevUserID).Error
	if err != nil {
		fmt.Println(3)
		res.FailWithMessage("接收人不存在", c)
		return
	}
	err = global.DB.Create(&model.MessageModel{
		SendUserID:       cr.SendUserID,
		SendUserNickName: sendUser.NickName,
		SendUserAvatar:   sendUser.Avatar,
		RevUserID:        cr.RevUserID,
		RevUserNickName:  revUser.NickName,
		RevUserAvatar:    revUser.Avatar,
		IsRead:           false,
		Content:          cr.Content,
	}).Error
	if err != nil {
		fmt.Println(4)
		res.FailWithMessage("消息发送失败", c)
		return
	}
	res.OkWithMessage("消息发送成功", c)
	return
}
