package repository

import (
	"sync"

	"github.com/johannes-kuhfuss/aoip-svc/config"
	"github.com/johannes-kuhfuss/aoip-svc/domain"
	"github.com/johannes-kuhfuss/services_utils/api_error"
)

type deviceList struct {
	mu   sync.RWMutex
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

func (drm *DeviceRepositoryMem) FindAll() (domain.Devices, int, api_error.ApiErr) {
	drm.DeviceList.mu.RLock()
	devices := drm.DeviceList.list
	drm.DeviceList.mu.Unlock()
	count := len(devices)
	return devices, count, nil
}
