package service

import (
	"go_code/gin-vue-blog/service/ImageService"
	"go_code/gin-vue-blog/service/RedisService"
)

type ServiceGroup struct {
	ImageService ImageService.ImageService
	RedisService RedisService.RedisService
}

var ServiceGroupApp = new(ServiceGroup)
