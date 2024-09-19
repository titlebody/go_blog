package service

import (
	"go_blog/service/image_ser"
	"go_blog/service/user_ser"
)

type ServiceGroup struct {
	ImageServiceGroup image_ser.ImageService
	UserService       user_ser.UserService
}

var ServiceApp = new(ServiceGroup)
