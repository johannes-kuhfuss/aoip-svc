package repository

import "github.com/johannes-kuhfuss/aoip-svc/config"

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
	return nil
}
