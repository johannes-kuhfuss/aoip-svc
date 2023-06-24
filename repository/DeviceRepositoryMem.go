package repository

import (
	"sync"

	"github.com/johannes-kuhfuss/aoip-svc/config"
	"github.com/johannes-kuhfuss/aoip-svc/domain"
	"github.com/johannes-kuhfuss/aoip-svc/dto"
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

func (drm *DeviceRepositoryMem) FindAll(safReq *dto.SortAndFilterRequest) (domain.Devices, int, api_error.ApiErr) {
	drm.DeviceList.mu.RLock()
	devices := drm.DeviceList.list
	drm.DeviceList.mu.RUnlock()
	devices = filter(devices, safReq)
	devices = sort(devices, safReq)
	devices = limit(devices, safReq)
	count := len(devices)
	return devices, count, nil
}

func filter(devs domain.Devices, safReq *dto.SortAndFilterRequest) domain.Devices {
	return devs
}

func sort(devs domain.Devices, safReq *dto.SortAndFilterRequest) domain.Devices {
	return devs
}

func limit(devs domain.Devices, safReq *dto.SortAndFilterRequest) domain.Devices {
	if len(devs) > safReq.Limit {
		return devs[0:safReq.Limit]
	} else {
		return devs
	}
}
