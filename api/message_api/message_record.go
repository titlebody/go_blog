package message_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/res"
	"go_blog/utils/jwts"
)

type MessageRecordRequest struct {
	UserID uint `json:"user_id" binding:"required" msg:"请输入查询用户的id"`
}

func (MessageApi) MessageRecordView(c *gin.Context) {
	var cr MessageRecordRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var _massageList []model.MessageModel
	massageList := make([]model.MessageModel, 0)
	// 获取最新的消息
	global.DB.Order("created_at asc").
		Find(&_massageList, "send_user_id = ? or rev_user_id = ?", claims.UserID, claims.UserID)

	for _, models := range _massageList {
		// 判断是一个组的条件
		// send_user_id 和rev_user_id 其中一个
		// 1 2 2 1
		// 1 3 3 1 是一组
		if models.RevUserID == cr.UserID || models.SendUserID == cr.UserID {
			massageList = append(massageList, models)

		}
	}
	//调用该接口双方后置为已读
	global.DB.Model(&model.MessageModel{}).Where("rev_user_id = ? and send_user_id = ?", claims.UserID, cr.UserID).
		Update("is_read", true)

	res.OkWithData(massageList, c)
}
