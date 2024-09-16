package menus_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/res"
)

func (MenuApi) MenuDetailView(c *gin.Context) {
	id := c.Param("id")
	var MenuModel model.MenuModel
	err := global.DB.Take(&MenuModel, id).Error
	if err != nil {
		res.FailWithMessage("菜单不存在", c)
		return
	}

	var menuBanner []model.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanner, "menu_id = ?", id)

	banners := make([]Banner, 0)
	for _, v := range menuBanner {
		if v.MenuID != MenuModel.ID {
			continue
		}
		banners = append(banners, Banner{
			ID:   v.BannerID,
			Path: v.BannerModel.Path,
		})
	}
	MenuResponse := MenuResponse{
		Banners:   banners,
		MenuModel: MenuModel,
	}
	res.OkWithData(MenuResponse, c)

}
