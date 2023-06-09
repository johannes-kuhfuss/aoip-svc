package service

import (
	"github.com/johannes-kuhfuss/aoip-svc/config"
	"github.com/johannes-kuhfuss/aoip-svc/domain"
	"github.com/johannes-kuhfuss/services_utils/logger"
	"github.com/micro/mdns"
)

type DiscoveryService interface {
	Discover()
}

type DefaultDiscoveryService struct {
	Cfg  *config.AppConfig
	repo domain.DiscoveryRepository
}

var (
	entriesCh    chan *mdns.ServiceEntry
	aoipServices = []string{"_netaudio-dbc._udp", "_netaudio-chan._udp", "_netaudio-cmc._udp"}
)

func NewDiscoveryService(cfg *config.AppConfig, repo domain.DiscoveryRepository) DefaultDiscoveryService {
	entriesCh = make(chan *mdns.ServiceEntry, 8)
	return DefaultDiscoveryService{
		Cfg:  cfg,
		repo: repo,
	}
}

func (s DefaultDiscoveryService) Discover() {
	logger.Info("Starting mDNS discovery...")
	s.Cfg.RunTime.MdnsQuery = true
	go s.storeEntries()
	err := mdns.Listen(entriesCh, nil)
	if err != nil {
		logger.Error("Error while listening for mDNS announcements", err)
	}
}

func (s DefaultDiscoveryService) storeEntries() {
	for entry := range entriesCh {
		/*
			logger.Info("*** Begin Info ***")
			logger.Info(fmt.Sprintf("entry.Addr: %v", entry.Addr))
			logger.Info(fmt.Sprintf("entry.AddrV4: %v", entry.AddrV4))
			logger.Info(fmt.Sprintf("entry.AddrV6: %v", entry.AddrV6))
			logger.Info(fmt.Sprintf("entry.Host: %v", entry.Host))
			logger.Info(fmt.Sprintf("entry.Info: %v", entry.Info))
			logger.Info(fmt.Sprintf("entry.InfoFields: %v", entry.InfoFields))
			logger.Info(fmt.Sprintf("entry.Name: %v", entry.Name))
			logger.Info(fmt.Sprintf("entry.Port: %v", entry.Port))
			logger.Info(fmt.Sprintf("entry.TTL: %v", entry.TTL))
			logger.Info("*** End Info ***")
		*/
		s.repo.Store(entry.Name, entry.Info)
	}
}
