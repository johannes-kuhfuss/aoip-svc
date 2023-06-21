package service

import (
	"encoding/json"
	"io/ioutil"

	"github.com/johannes-kuhfuss/aoip-svc/config"
	"github.com/johannes-kuhfuss/aoip-svc/domain"
	"github.com/johannes-kuhfuss/services_utils/logger"
)

type DeviceService interface {
	Run()
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

func (s DefaultDeviceService) Run() {
	devs := []domain.Device{}
	sampleData, err := ioutil.ReadFile("../sample-data/coloRadio-dante-devices.json")
	if err != nil {
		logger.Error("Oops: ", err)
	}
	err = json.Unmarshal(sampleData, &devs)
	if err != nil {
		logger.Error("Oops: ", err)
	}
}
