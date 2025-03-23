package service

import "go_code/gin-vue-blog/service/ImageService"

type ServiceGroup struct {
	ImageService ImageService.ImageService
}

var ServiceGroupApp = new(ServiceGroup)
