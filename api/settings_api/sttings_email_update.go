package settings_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/config"
	"go_blog/core"
	"go_blog/global"
	"go_blog/model/res"
)

func (SettingsApi) SettingsEmailUpdateView(c *gin.Context) {
	var cr config.Email
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	global.Config.Email = cr
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OKWith(c)
}
