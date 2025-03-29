package api

import (
	"go_code/gin-vue-blog/api/adApi"
	"go_code/gin-vue-blog/api/imagesApi"
	"go_code/gin-vue-blog/api/menuApi"
	"go_code/gin-vue-blog/api/messageApi"
	"go_code/gin-vue-blog/api/settingsApi"
	"go_code/gin-vue-blog/api/tagApi"
	"go_code/gin-vue-blog/api/userApi"
)

type ApiGroup struct {
	SettingsApi settingsApi.SettingsApi
	ImagesApi   imagesApi.ImagesApi
	AdApi       adApi.AdApi
	MenuApi     menuApi.MenuApi
	UserApi     userApi.UserApi
	TagApi      tagApi.TagApi
	MessageApi  messageApi.MessageApi
}

var ApiGroupApp = new(ApiGroup)
