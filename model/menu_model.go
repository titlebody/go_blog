package model

import "go_blog/model/c_type"

// MenuModel 菜单
type MenuModel struct {
	MODEL
	Title        string        `gorm:"size:32" json:"title"`                                                                           // 菜单标题
	Path         string        `gorm:"size:32" json:"path"`                                                                            // 菜单标题
	Slogan       string        `gorm:"size:64" json:"slogan"`                                                                          // 菜单标语
	Abstract     c_type.Array  `gorm:"type:string" json:"abstract"`                                                                    // 简介
	AbstractTime int           `json:"abstract_time"`                                                                                  // 简介切换时间
	Banners      []BannerModel `gorm:"many2many:menu_banner_model;joinForeignKey:MenuID;JoinReferences:BannerID" json:"banner_images"` // 菜单图片
	BannerTime   int           `json:"banner_time"`                                                                                    // 菜单图片切换时间
	Sort         int           `gorm:"size:10" json:"sort"`                                                                            // 菜单顺序
}
