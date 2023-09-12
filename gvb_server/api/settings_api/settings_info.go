package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models/res"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	//c.JSON(http.StatusOK, gin.H{
	//	"Msg": "hhh",
	//})
	res.Fail(1, map[string]any{}, "OK", c)
}
