package settingsApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/models/common"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	common.OK(map[string]string{}, "xxx", c)
}
