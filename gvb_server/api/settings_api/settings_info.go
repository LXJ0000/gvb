package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	switch cr.Name {
	case "email":
		res.OKWithData(global.Config.Email, c)
	case "qq":
		res.OKWithData(global.Config.QQ, c)
	case "qiniu":
		res.OKWithData(global.Config.QiNiu, c)
	case "jwt":
		res.OKWithData(global.Config.Jwt, c)
	case "site":
		res.OKWithData(global.Config.SiteInfo, c)
	default:
		res.FailWithMessage("配置信息不存在", c)
	}
}
