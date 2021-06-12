package postgres

import (
	"database/sql"
	"github.com/carrot-systems/carrot-automation/src/core/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type RunningHistory struct {
	gorm.Model
	ID      uuid.UUID `gorm:"type:uuid;primary_key"`
	Started time.Time
	Done    sql.NullTime
	Status  int
}

type ActionVariable struct {
	gorm.Model
	ID    uuid.UUID `gorm:"type:uuid;primary_key"`
	Name  string
	Value string
}

type Action struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	ServiceId string
	ActionId  string
	Trigger   bool
}

type Workflow struct {
	gorm.Model
	ID      uuid.UUID `gorm:"type:uuid;primary_key"`
	Name    string
	User    *uuid.UUID
	History []*uuid.UUID `gorm:"type:text"`
	Trigger *uuid.UUID
}

type workflowRepo struct {
	db *gorm.DB
}

func historyFromDomain(history *domain.RunningHistory) *RunningHistory {
	return &RunningHistory{
		ID:      history.ID,
		Started: time.Time{},
		Done:    sql.NullTime{},
		Status:  0,
	}
}

func actionFromDomain(action *domain.Action) *Action {
	return &Action{
		ID:       action.ID,
		ActionId: action.ActionId,
		Trigger:  action.IsTrigger,
	}
}

func workflowToDomain(workflow *Workflow) *domain.Workflow {
	return &domain.Workflow{
		ID:      workflow.ID,
		Name:    workflow.Name,
		User:    nil,
		History: nil,
		Trigger: nil,
	}
}

func workflowsToDomain(workflows []*Workflow) []*domain.Workflow {
	domainWorkflows := []*domain.Workflow{}

	for _, workflow := range workflows {
		domainWorkflows = append(domainWorkflows, workflowToDomain(workflow))
	}

	return domainWorkflows
}

func workflowFromDomain(workflow *domain.Workflow) *Workflow {
	var history []*uuid.UUID

	for _, entry := range workflow.History {
		history = append(history, entry)
	}

	var trigger *uuid.UUID
	trigger = nil

	if workflow.Trigger != nil {
		trigger = workflow.Trigger
	}

	return &Workflow{
		ID:      workflow.ID,
		Name:    workflow.Name,
		User:    workflow.User,
		History: history,
		Trigger: trigger,
	}
}

func (w workflowRepo) CreateWorkflow(workflow *domain.Workflow) error {
	workflowToCreate := workflowFromDomain(workflow)

	result := w.db.Create(workflowToCreate)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (w workflowRepo) FindByName(user uuid.UUID, name string) (*domain.Workflow, error) {
	var workflow *Workflow

	query := w.db.Where("\"user\" = ? AND name = ?", user, name).First(&workflow)

	if query.Error != nil {
		return nil, query.Error
	}

	return workflowToDomain(workflow), nil
}

func (w workflowRepo) FindByUser(user uuid.UUID) ([]*domain.Workflow, error) {
	var workflows []*Workflow

	query := w.db.Where("\"user\" = ?", user).Find(&workflows)

	if query.Error != nil {
		return nil, query.Error
	}

	return workflowsToDomain(workflows), nil
}

func (w workflowRepo) FindById(id string) (*domain.Workflow, error) {
	var workflow *Workflow

	query := w.db.Where("\"id\" = ?", id).First(&workflow)

	if query.Error != nil {
		return nil, query.Error
	}

	return workflowToDomain(workflow), nil
}

func (w workflowRepo) DeleteWorkflow(workflow *domain.Workflow) error {
	idToRemove := workflow.ID

	query := w.db.Where("id = ?", idToRemove).Delete(&Workflow{})

	return query.Error
}

func NewWorkflowRepo(db *gorm.DB) workflowRepo {
	return workflowRepo{
		db: db,
	}
}
