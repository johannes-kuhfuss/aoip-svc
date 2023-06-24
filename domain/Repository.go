package domain

import (
	"github.com/johannes-kuhfuss/aoip-svc/dto"
	"github.com/johannes-kuhfuss/services_utils/api_error"
)

type DeviceRepository interface {
	Store(Devices)
	FindAll(*dto.SortAndFilterRequest) (Devices, int, api_error.ApiErr)
}
