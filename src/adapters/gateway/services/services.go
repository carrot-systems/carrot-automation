package services

import (
	"github.com/carrot-systems/carrot-automation/src/adapters/gateway/services/onesignal"
	"github.com/carrot-systems/carrot-automation/src/adapters/gateway/services/standardtools"
	"github.com/carrot-systems/carrot-automation/src/core/domain"
)

type ServiceManager struct {
	services []domain.Service
}

func CreateServiceManager() ServiceManager {
	return ServiceManager{
		services: []domain.Service{
			onesignal.NewOnesignal(),
			standardtools.NewStandardTools(),
		},
	}
}

func (sm ServiceManager) GetServices() []domain.Service {
	return sm.services
}

func (sm ServiceManager) GetService(id string) (*domain.Service, error) {
	for _, service := range sm.services {
		if service.Id == id {
			return &service, nil
		}
	}
	return nil, domain.ErrServiceNotFound
}
