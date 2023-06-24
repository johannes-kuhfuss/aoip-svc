package repository

import (
	"fmt"
	"sync"

	"github.com/johannes-kuhfuss/aoip-svc/config"
	"github.com/johannes-kuhfuss/aoip-svc/domain"
	"github.com/johannes-kuhfuss/aoip-svc/dto"
	"github.com/johannes-kuhfuss/services_utils/api_error"
	"github.com/johannes-kuhfuss/services_utils/logger"
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
	allDevices := drm.DeviceList.list
	drm.DeviceList.mu.RUnlock()
	// sort and filter
	logger.Info(fmt.Sprintf("saf: %#v", safReq))
	count := len(allDevices)
	return allDevices, count, nil
}
