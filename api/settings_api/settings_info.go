package settings_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model/res"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.OkWithData(global.Config.SiteInfo, c)
}
