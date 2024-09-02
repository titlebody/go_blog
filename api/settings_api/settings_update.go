package settings_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/config"
	"go_blog/core"
	"go_blog/global"
	"go_blog/model/res"
)

// SettingsUpdateView 修改某一项配置信息
func (SettingsApi) SettingsUpdateView(c *gin.Context) {

	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	switch cr.Name {
	case "site":
		var info config.SiteInfo
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithMessage(res.ArgumentError, c)
			return
		}
		global.Config.SiteInfo = info
	case "email":
		var info config.Email
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithMessage(res.ArgumentError, c)
			return
		}
		global.Config.Email = info
	case "qq":
		var info config.QQ
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithMessage(res.ArgumentError, c)
			return
		}
		global.Config.QQ = info
	case "qiniu":
		var info config.QiNiu
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithMessage(res.ArgumentError, c)
			return
		}
		global.Config.QiNiu = info
	case "jwt":
		var info config.JWT
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithMessage(res.ArgumentError, c)
			return
		}
		global.Config.JWT = info
	default:
		res.FailWithMessage("没有对应的配置信息", c)
	}
	err = core.SetYaml()
	if err != nil {
		res.FailWithCode(res.SettingsError, c)
		return
	}
	res.OKWith(c)

}
