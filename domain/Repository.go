package domain

import "github.com/johannes-kuhfuss/services_utils/api_error"

type DeviceRepository interface {
	Store(Devices)
	FindAll() (Devices, int, api_error.ApiErr)
}
