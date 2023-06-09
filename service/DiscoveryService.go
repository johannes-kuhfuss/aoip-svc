package service

import (
	"fmt"

	"github.com/johannes-kuhfuss/aoip-svc/config"
	"github.com/johannes-kuhfuss/services_utils/logger"
	"github.com/micro/mdns"
)

type DiscoveryService interface {
	Discover()
}

type DefaultDiscoveryService struct {
	Cfg *config.AppConfig
}

var (
	entriesCh chan *mdns.ServiceEntry
)

func NewDiscoveryService(cfg *config.AppConfig) DefaultDiscoveryService {
	entriesCh = make(chan *mdns.ServiceEntry, 8)
	return DefaultDiscoveryService{
		Cfg: cfg,
	}
}

func (s DefaultDiscoveryService) Discover() {
	logger.Info("Starting mDNS discovery...")
	go listEntries()
	err := mdns.Listen(entriesCh, nil)
	if err != nil {
		logger.Error("Error while listening for mDNS announcements", err)
	}
}

func listEntries() {
	for entry := range entriesCh {
		msg := fmt.Sprintf("Got new entry: %v", entry)
		logger.Info(msg)
	}
}
