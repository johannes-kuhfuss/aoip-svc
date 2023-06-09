package domain

type DiscoveryRepository interface {
	Store(string, string) error
}
