package repository

import (
	"fmt"

	"github.com/johannes-kuhfuss/aoip-svc/config"
	"github.com/johannes-kuhfuss/services_utils/logger"
)

type DiscoveryRepositoryMem struct {
	Cfg *config.AppConfig
}

var (
	services = map[string]string{}
)

func NewDiscoveryRepository(cfg *config.AppConfig) DiscoveryRepositoryMem {
	return DiscoveryRepositoryMem{
		Cfg: cfg,
	}
}

func (r DiscoveryRepositoryMem) Store(svc, data string) (err error) {
	services[svc] = data
	logger.Info(fmt.Sprintf("Stored %v with value %v", svc, data))
	logger.Info(fmt.Sprintf("Length of list: %v", len(services)))
	return nil
}
