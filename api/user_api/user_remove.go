package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/res"
	"gorm.io/gorm"
)

func (UserAPI) UserRemoveView(c *gin.Context) {
	var cr model.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var userList []model.UserModel
	count := global.DB.Find(&userList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("用户不存在", c)
		return
	}

	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// 删除用户登录数据
		if err := tx.Where("user_id IN ?", cr.IDList).Delete(&model.LoginDataModel{}).Error; err != nil {
			global.Log.Error(err)
			return err
		}

		// 删除用户收藏的文章
		if err := tx.Where("user_id IN ?", cr.IDList).Delete(&model.UserCollectModel{}).Error; err != nil {
			global.Log.Error(err)
			return err
		}

		// 删除用户
		if err := tx.Delete(&userList).Error; err != nil {
			global.Log.Error(err)
			return err
		}

		return nil
	})

	if err != nil {
		res.FailWithMessage("删除失败", c)
		return
	}

	res.OkWithMessage(fmt.Sprintf("共删除%d个用户", count), c)
}
