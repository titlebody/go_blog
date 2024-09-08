package main

import (
	"fmt"
	"github.com/fatih/structs"
	"go_blog/model"
)

type AdvertRequest struct {
	model.MODEL `structs:"-"`
	Title       string `json:"title" binding:"required" msg:"请输入标题" structs:"title"` // 广告标题
	Href        string `json:"href" binding:"required,url" msg:"广告链接非法" structs:"-"` // 广告链接
	Images      string `json:"images" binding:"required,url" msg:"图片地址非法"`           // 广告图片
	IsShow      bool   `json:"is_show" binding:"required" msg:"请选择是否显示"`             // 是否显示
}

func main() {
	u1 := AdvertRequest{
		Title:  "title",
		Href:   "http://www.baidu.com",
		Images: "http://www.baidu.com",
		IsShow: true,
	}
	m3 := structs.Map(&u1)
	for k, v := range m3 {
		fmt.Printf("key:%v value:%v value type:%T\n", k, v, v)
	}
	fmt.Println(m3)
}
