package settingsApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/config"
	"go_code/gin-vue-blog/core"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models/res"
)

// SettingsInfoUpdateView 修改某一项配置信息
func (settingsApi *SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var cr SettingsInfoUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "site":
		var info config.SiteInfo
		err = c.ShouldBind(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
		}
		global.Config.SiteInfo = info
	case "email":
		var info config.Email
		err = c.ShouldBind(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
		}
		global.Config.Email = info
	case "qq":
		var info config.QQ
		err = c.ShouldBind(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
		}
		global.Config.QQ = info
	case "jwt":
		var info config.Jwt
		err = c.ShouldBind(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
		}
		global.Config.Jwt = info
	default:
		res.FailWithMsg("没有对应的配置信息", c)
		return
	}
	core.SetYaml()
	res.OKWithNothing(c)
}
