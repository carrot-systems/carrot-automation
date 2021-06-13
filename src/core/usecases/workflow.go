package usecases

import "github.com/carrot-systems/carrot-automation/src/core/domain"

func (i interactor) CreateWorkflow(user *domain.User, name string) error {
	workflow := &domain.Workflow{
		Name:    name,
		User:    &user.UserID,
		History: nil,
		Trigger: nil,
	}

	workflow.Initialize()

	workflowSameName, err := i.workflowRepo.FindByName(user.UserID, name)

	if err == nil && workflowSameName != nil {
		//A workflow with the same name already exists for this user
		return domain.ErrWorkflowAlreadyExistingWithThisName
	}

	err = i.workflowRepo.CreateWorkflow(workflow)

	if err != nil {
		return err
	}

	return nil
}

func (i interactor) GetUserWorkflows(user *domain.User) ([]*domain.Workflow, error) {
	workflows, err := i.workflowRepo.FindByUser(user.UserID)

	if err != nil {
		return nil, err
	}

	return workflows, nil
}

func (i interactor) DeleteWorkflow(user *domain.User, id string) error {
	//TODO: check if workflow belong to user
	workflow, err := i.workflowRepo.FindById(id)

	if err != nil || workflow == nil {
		return domain.ErrWorkflowNotFound
	}

	i.workflowRepo.DeleteWorkflow(workflow)

	return nil
}

func (i interactor) SetWorkflowData(user *domain.User, workflowId string, workflowData domain.WorkflowContent) error {
	return domain.ErrFeatureNotImplemented
}
