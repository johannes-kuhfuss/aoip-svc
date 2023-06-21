package service

import (
	"github.com/johannes-kuhfuss/aoip-svc/config"
)

type DeviceService interface {
	Discover()
}

type DefaultDeviceService struct {
	Cfg *config.AppConfig
}

var ()

func NewDeviceService(cfg *config.AppConfig) DefaultDeviceService {
	return DefaultDeviceService{
		Cfg: cfg,
	}
}
