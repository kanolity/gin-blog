package settingsApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models/res"
)

type SettingsInfoUri struct {
	Name string `uri:"name"`
}

// SettingsInfoView 展示某一项配置信息
// @Tags         设置管理
// @Summary      展示某一项配置信息
// @Description  展示某一项配置信息
// @Param name path string  true "设置名"
// @Router       /api/settings/{name} [get]
// @Produce json
// @Success      200  {object}  res.Resp{data=string}
func (settingsApi *SettingsApi) SettingsInfoView(c *gin.Context) {
	var cr SettingsInfoUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "site":
		res.OKWithData(global.Config.SiteInfo, c)
	case "email":
		res.OKWithData(global.Config.Email, c)
	case "qq":
		res.OKWithData(global.Config.QQ, c)
	case "jwt":
		res.OKWithData(global.Config.Jwt, c)
	default:
		res.FailWithMsg("没有对应的配置信息", c)
	}
}
