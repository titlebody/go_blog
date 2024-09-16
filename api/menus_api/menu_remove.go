package menus_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/res"
	"gorm.io/gorm"
)

func (MenuApi) MenuRemoveView(c *gin.Context) {
	var cr model.RemoveRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var menuModel []model.MenuModel
	count := global.DB.Find(&menuModel, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("菜单不存在", c)
		return
	}

	err := global.DB.Transaction(func(tx *gorm.DB) error {
		//删除第三张表
		err := global.DB.Model(&menuModel).Association("Banners").Clear()
		if err != nil {
			return err
		}

		err1 := global.DB.Delete(&menuModel).Error
		if err1 != nil {
			global.Log.Error(err1)
			return err1
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除菜单失败！", c)
		return
	}

	res.OkWithMessage(fmt.Sprintf("删除成功%d个菜单", count), c)

}
