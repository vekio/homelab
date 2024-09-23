package homelab

import "sync"

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
