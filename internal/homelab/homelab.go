package homelab

import (
	"sync"

	"github.com/vekio/homelab/internal/config"
)

type Homelab struct {
	Services Services
}

func NewHomelabApp(conf *config.ConfigManager[config.Config]) (Homelab, error) {
	var services Services = make(Services)
	var wg sync.WaitGroup
	errCh := make(chan error, len(conf.Data.Services))

	for serviceName, serviceConfig := range conf.Data.Services {
		wg.Add(1)
		go func(name, context string, composeFiles []string, conf *config.ConfigManager[config.Config]) {
			defer wg.Done()
			service, err := NewService(name, context, composeFiles, conf)
			if err != nil {
				errCh <- err
				return
			}
			services[name] = service
		}(serviceName, serviceConfig.Context, serviceConfig.ComposeFiles, conf)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			return Homelab{}, err
		}
	}

	homelab := Homelab{
		Services: services,
	}

	return homelab, nil
}
