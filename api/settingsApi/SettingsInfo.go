package settingsApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models/common"
)

type SettingsInfoUri struct {
	Name string `uri:"name"`
}

// SettingsInfoView 展示某一项配置信息
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	var cr SettingsInfoUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		common.FailWithCode(common.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "site":
		common.OKWithData(global.Config.SiteInfo, c)
	case "email":
		common.OKWithData(global.Config.Email, c)
	case "qq":
		common.OKWithData(global.Config.QQ, c)
	case "jwt":
		common.OKWithData(global.Config.Jwt, c)
	default:
		common.FailWithMsg("没有对应的配置信息", c)
	}
}
