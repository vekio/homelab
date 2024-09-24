package homelab

import (
	"fmt"
	"sync"
)

type Homelab struct {
	Services Services
}

func NewHomelab() (Homelab, error) {
	var services Services = make(Services)
	var wg sync.WaitGroup
	errCh := make(chan error, len(settings.Services))

	for serviceName, serviceConfig := range settings.Services {
		wg.Add(1)
		go func(name, context string, composeFiles []string) {
			defer wg.Done()
			service, err := NewService(name, context, composeFiles)
			if err != nil {
				errCh <- err
				return
			}
			services[name] = service
		}(serviceName, serviceConfig.Context, serviceConfig.ComposeFiles)
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

func (h *Homelab) ServicesNames() []string {
	keys := make([]string, 0, len(h.Services))
	for k := range h.Services {
		keys = append(keys, k)
	}
	return keys
}

func (h *Homelab) ServiceByName(name string) (*Service, error) {
	service, ok := h.Services[name]
	if !ok {
		return nil, fmt.Errorf("service %s not found", name)
	}
	return service, nil
}
