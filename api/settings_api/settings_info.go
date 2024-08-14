package settings_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/model/res"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.FailWithCode(res.SettingsError, c)
}
