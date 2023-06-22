package service

import (
	"encoding/json"
	"fmt"
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
	var devices []domain.Device
	var objMap map[string]json.RawMessage
	var dev domain.Device
	sampleData, err := ioutil.ReadFile("./sample-data/coloRadio-dante-devices.json")
	if err != nil {
		logger.Error("Oops: ", err)
	}
	err = json.Unmarshal(sampleData, &objMap)
	if err != nil {
		logger.Error("Oops: ", err)
	}
	for obj := range objMap {
		logger.Info(fmt.Sprintf("Object: %v", obj))
		rawDev := domain.RawDevice{}
		err = json.Unmarshal(objMap[obj], &rawDev)
		if err != nil {
			logger.Error("Oops: ", err)
		}
		dev, err = dev.FromRawDevice(rawDev)
		if err != nil {
			logger.Error("Oops: ", err)
		}
		devices = append(devices, dev)
	}
}
