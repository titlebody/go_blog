package menus_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/res"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type MenuResponse struct {
	model.MenuModel
	Banners []Banner `json:"banners"`
}

func (MenuApi) MenuListView(c *gin.Context) {
	// 查询菜单列表
	var menuList []model.MenuModel
	var menuIDList []uint
	global.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIDList)
	// 查连接表
	var menuBanners []model.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id in (?)", menuIDList)
	var menus []MenuResponse
	for _, menu := range menuList {
		var banners = make([]Banner, 0)
		for _, menuBanner := range menuBanners {
			if menuBanner.MenuID == menu.ID {
				banners = append(banners, Banner{
					ID:   menuBanner.BannerID,
					Path: menuBanner.BannerModel.Path,
				})
			}
		}
		menus = append(menus, MenuResponse{
			MenuModel: menu,
			Banners:   banners,
		})
	}

	res.OkWithData(menus, c)
	return
}
