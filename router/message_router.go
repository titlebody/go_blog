package router

import (
	"go_blog/api"
	"go_blog/middleware"
)

func (r RouterGroup) MessageRouter() {
	api := api.ApiGroupApp.MessageApi

	r.POST("/messages", api.MessageCreateView)                             //发送消息
	r.GET("/messages_all", api.MessageListAllView)                         // 管理员查看所有消息
	r.GET("/messages_list", middleware.JwtAuth(), api.MessageListView)     // 查看自己的消息列表
	r.GET("/messages_record", middleware.JwtAuth(), api.MessageRecordView) // 查看自己与某用户的消息记录1对1
}
