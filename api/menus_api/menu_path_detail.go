package menus_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/res"
)

type MenuDetailRequest struct {
	Path string `form:"path"`
}

func (MenuApi) MenuPathDetailView(c *gin.Context) {
	var cr MenuDetailRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithMessage("参数错误", c)
		return
	}

	var MenuModel model.MenuModel
	err = global.DB.Take(&MenuModel, "path = ?", cr.Path).Error
	if err != nil {
		res.FailWithMessage("菜单不存在", c)
		return
	}

	var menuBanner []model.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanner, "menu_id = ?", MenuModel.ID)

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
