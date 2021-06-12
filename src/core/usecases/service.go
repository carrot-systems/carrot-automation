package usecases

import "github.com/carrot-systems/carrot-automation/src/core/domain"

func (i interactor) GetAllServices() ([]domain.Service, error) {
	return i.serviceGateway.GetServices(), nil
}

func (i interactor) GetServiceActions(serviceName string) ([]domain.Action, error) {
	service, err := i.serviceGateway.GetService(serviceName)

	if err != nil {
		return nil, err
	}

	return service.GetActions(), nil
}
