package usecases

import (
	"github.com/carrot-systems/carrot-automation/src/adapters/gateway/jwt"
	"github.com/carrot-systems/carrot-automation/src/core/domain"
	"github.com/google/uuid"
)

type Logger interface {
	Error(err error)
	Fatal(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
}

type ServiceGateway interface {
	GetServices() []domain.Service
	GetService(service string) (*domain.Service, error)
}

type WorkflowRepo interface {
	CreateWorkflow(workflow *domain.Workflow) error
	FindByName(user uuid.UUID, name string) (*domain.Workflow, error)
	FindByUser(user uuid.UUID) ([]*domain.Workflow, error)
	FindById(id string) (*domain.Workflow, error)
	DeleteWorkflow(workflow *domain.Workflow) error
}

type interactor struct {
	serviceGateway ServiceGateway
	workflowRepo   WorkflowRepo
	jwtInstance    *jwt.JwtInstance
}

func NewInteractor(jwt *jwt.JwtInstance, sG ServiceGateway, wR WorkflowRepo) interactor {
	return interactor{
		serviceGateway: sG,
		workflowRepo:   wR,
		jwtInstance:    jwt,
	}
}
