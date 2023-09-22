package service

import (
	"gvb_server/service/image_service"
	"gvb_server/service/redis_service"
	"gvb_server/service/user_service"
)

type ServiceGroup struct {
	ImageService image_service.ImageService
	USerService  user_service.UserService
	RedisService redis_service.RedisService
}

var ServiceApp = new(ServiceGroup)
