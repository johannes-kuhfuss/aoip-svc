package repository

import (
	"sync"

	"github.com/johannes-kuhfuss/aoip-svc/config"
	"github.com/johannes-kuhfuss/aoip-svc/domain"
)

type deviceList struct {
	mu   sync.Mutex
	list domain.Devices
}

type DeviceRepositoryMem struct {
	Cfg        *config.AppConfig
	DeviceList deviceList
}

func NewDeviceRepositoryMem(cfg *config.AppConfig) DeviceRepositoryMem {
	return DeviceRepositoryMem{
		Cfg:        cfg,
		DeviceList: deviceList{},
	}
}

func (drm *DeviceRepositoryMem) Store(list domain.Devices) {
	drm.DeviceList.mu.Lock()
	drm.DeviceList.list = list
	drm.DeviceList.mu.Unlock()
}
