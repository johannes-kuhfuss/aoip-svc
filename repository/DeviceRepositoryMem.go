package repository

import (
	"github.com/johannes-kuhfuss/aoip-svc/config"
)

type DeviceRepositoryMem struct {
	Cfg *config.AppConfig
}

func NewDeviceRepository(cfg *config.AppConfig) DeviceRepositoryMem {
	return DeviceRepositoryMem{
		Cfg: cfg,
	}
}
