package service

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/johannes-kuhfuss/aoip-svc/config"
	"github.com/johannes-kuhfuss/aoip-svc/domain"
	"github.com/johannes-kuhfuss/services_utils/logger"
)

type DeviceService interface {
	Run()
}

type DefaultDeviceService struct {
	Cfg  *config.AppConfig
	Repo domain.DeviceRepository
}

var ()

func NewDeviceService(cfg *config.AppConfig, repo domain.DeviceRepository) DefaultDeviceService {
	return DefaultDeviceService{
		Cfg:  cfg,
		Repo: repo,
	}
}

func (s DefaultDeviceService) Run() {
	for s.Cfg.RunTime.RunDiscover == true {
		s.Discover()
		time.Sleep(time.Duration(s.Cfg.DeviceDiscovery.IntervalSec) * time.Second)
	}
}

func (s DefaultDeviceService) Discover() {
	var (
		devices domain.Devices
		rawData map[string]json.RawMessage
		dev     domain.Device
	)
	logger.Info("Start new discovery cycle...")
	data, err := retrieveData()
	err = json.Unmarshal(data, &rawData)
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
	s.Repo.Store(devices)
	logger.Info("End of new discovery cycle.")
}

func retrieveData() ([]byte, error) {
	data, err := ioutil.ReadFile("./sample-data/coloRadio-dante-devices.json")
	if err != nil {
		logger.Error("Error reading sample file: ", err)
		return nil, err
	}
	return data, nil
}
