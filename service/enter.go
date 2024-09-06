package service

import (
	"go_blog/service/image_ser"
)

type ServiceGroup struct {
	ImageServiceGroup image_ser.ImageService
}

var ServiceApp = new(ServiceGroup)
