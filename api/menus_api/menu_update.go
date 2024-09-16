package menus_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/res"
)

func (MenuApi) MenuUpdateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	id := c.Param("id")
	//先把banner清空
	var menuModel model.MenuModel
	err = global.DB.Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMessage("菜单不存在", c)
		return
	}
	// 先清空
	err = global.DB.Model(&menuModel).Association("Banners").Clear()
	if err != nil {
		res.FailWithMessage("数据清空失败", c)
		return
	}

	// 如果选择了banner,那就添加
	if len(cr.ImagesSortList) > 0 {
		var bannerList []model.MenuBannerModel
		for _, v := range cr.ImagesSortList {
			bannerList = append(bannerList, model.MenuBannerModel{
				MenuID:   menuModel.ID,
				BannerID: v.ImageID,
				Sort:     v.Sort,
			})
		}
		//global.DB.Model(&menuModel).Association("Banners").Replace(menuModel)
		err := global.DB.Create(&bannerList).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("修改菜单图片失败", c)
			return
		}
	}
	// 普通更新
	maps := structs.Map(&cr)
	err = global.DB.Model(&menuModel).Updates(maps).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改菜单失败", c)
		return
	}
	res.OkWithMessage("修改菜单成功", c)
}
