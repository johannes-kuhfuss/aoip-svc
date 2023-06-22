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
	var (
		devices domain.Devices
		rawData map[string]json.RawMessage
		dev     domain.Device
	)
	sampleData, err := ioutil.ReadFile("./sample-data/coloRadio-dante-devices.json")
	if err != nil {
		logger.Error("Error reading sample file: ", err)
	}
	err = json.Unmarshal(sampleData, &rawData)
	if err != nil {
		logger.Error("Error converting output from netaudio: ", err)
	}
	for item := range rawData {
		rawDev := domain.RawDevice{}
		err = json.Unmarshal(rawData[item], &rawDev)
		if err != nil {
			logger.Error("Error converting data into raw device data: ", err)
			return
		}
		dev, err = dev.FromRawDevice(rawDev)
		if err != nil {
			logger.Error("Error converting raw device into device: ", err)
			return
		}
		devices = append(devices, dev)
	}
}
