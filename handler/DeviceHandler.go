package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johannes-kuhfuss/aoip-svc/config"
	"github.com/johannes-kuhfuss/aoip-svc/service"
	"github.com/johannes-kuhfuss/services_utils/logger"
)

type DeviceHandler struct {
	Cfg     *config.AppConfig
	Service service.DeviceService
}

func NewDeviceHandler(cfg *config.AppConfig, svc service.DeviceService) DeviceHandler {
	return DeviceHandler{
		Cfg:     cfg,
		Service: svc,
	}
}

func (dh *DeviceHandler) GetAllDevices(c *gin.Context) {
	jobs, totalCount, err := dh.Service.GetAllDevices()
	if err != nil {
		logger.Error("Service error while getting all jobs", err)
		c.JSON(err.StatusCode(), err)
		return
	}
	countStr := fmt.Sprintf("%v", totalCount)
	c.Header("X-Total-Count", countStr)
	c.JSON(http.StatusOK, jobs)
}
