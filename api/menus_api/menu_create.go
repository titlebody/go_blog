package menus_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/c_type"
	"go_blog/model/res"
)

type ImageStore struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenuRequest struct {
	Title          string       `json:"title" binding:"required" msg:"请完善菜单名称" structs:"title"`
	Path           string       `json:"path" binding:"required" msg:"请完善菜单路径" structs:"path"`
	Slogan         string       `json:"slogan" structs:"slogan"`
	Abstract       c_type.Array `json:"abstract" structs:"abstract"`
	AbstractTime   int          `json:"abstract_time" structs:"abstract_time"`                // 切换的时间
	BannerTime     int          `json:"banner_time" structs:"banner_time"`                    //切换的时间
	Sort           int          `json:"sort" binding:"required" msg:"请完善菜单序号" structs:"sort"` // 菜单序号
	ImagesSortList []ImageStore `json:"images_sort_list" structs:"-"`                         // 具体图片的排序
}

func (MenuApi) MenuCreateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//重复值判断
	var menuList []model.MenuModel
	count := global.DB.Find(&menuList, "path=? or title=?", cr.Path, cr.Title).RowsAffected
	if count > 0 {
		res.FailWithMessage("菜单名称或路径重复", c)
		return
	}

	menuModel := model.MenuModel{
		Title:        cr.Title,
		Path:         cr.Path,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}

	err = global.DB.Create(&menuModel).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("菜单添加失败", c)
		return
	}

	if len(cr.ImagesSortList) == 0 {
		res.OkWithMessage("添加菜单成功", c)
		return
	}

	var menuBannerList []model.MenuBannerModel

	for _, v := range cr.ImagesSortList {
		menuBannerList = append(menuBannerList, model.MenuBannerModel{
			MenuID:   menuModel.ID,
			BannerID: v.ImageID,
			Sort:     v.Sort,
		})

	}

	//第三张表入库
	err = global.DB.Create(menuBannerList).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("菜单图片添加失败", c)
		return
	}
	res.OkWithMessage("添加菜单成功", c)
}
