package model

import "go_blog/model/c_type"

type MenuModel struct {
	MODEL
	MenuTitle    string        `gorm:"size:32" json:"menu_title"`                                                                      // 菜单标题
	MenuTitleEn  string        `gorm:"size:32" json:"menu_title_en"`                                                                   // 菜单标题
	Slogan       string        `gorm:"size:64" json:"slogan"`                                                                          // 菜单标语
	Abstract     c_type.Array  `gorm:"type:string" json:"abstract"`                                                                    // 简介
	AbstractTime int           `json:"abstract_time"`                                                                                  // 简介切换时间
	Banners      []BannerModel `gorm:"many2many:menu_banner_model;joinForeignKey:MenuID;JoinReferences:BannerID" json:"banner_images"` // 菜单图片
	BannerTime   int           `json:"banner_time"`                                                                                    // 菜单图片切换时间
	Sort         int           `gorm:"size:10" json:"sort"`                                                                            // 菜单顺序
}
