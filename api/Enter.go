package api

import (
	"go_code/gin-vue-blog/api/imagesApi"
	"go_code/gin-vue-blog/api/settingsApi"
)

type ApiGroup struct {
	SettingsApi settingsApi.SettingsApi
	ImagesApi   imagesApi.ImagesApi
}

var ApiGroupApp = new(ApiGroup)
