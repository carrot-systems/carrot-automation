package usecases

import (
	"github.com/carrot-systems/carrot-automation/src/core/domain"
)

type Usecases interface {
	//Workflows
	CreateWorkflow(user *domain.User, name string) error
	GetUserWorkflows(user *domain.User) ([]*domain.Workflow, error)
	DeleteWorkflow(user *domain.User, id string) error

	//Services
	GetAllServices() ([]domain.Service, error)
	GetServiceActions(service string) ([]domain.Action, error)
}
