package usecases

import (
	"github.com/carrot-systems/carrot-automation/src/core/domain"
)

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
	workflow, err := i.workflowRepo.FindById(user.UserID, id)

	if err != nil || workflow == nil {
		return domain.ErrWorkflowNotFound
	}

	if *workflow.User != user.UserID {
		return domain.ErrWorkflowNotFound
	}

	i.workflowRepo.DeleteWorkflow(workflow)

	return nil
}

func (i interactor) findTriggerInWorkflow(workflowData domain.WorkflowContent) (*domain.WorkflowAction, error) {
	var action *domain.WorkflowAction = nil

	for _, currentAction := range workflowData.Actions {
		if currentAction.Trigger {
			if action != nil {
				return nil, domain.ErrWorkflowHasMoreThanOneTrigger
			}
			action = &currentAction
		}
	}

	if action == nil {
		return nil, domain.ErrWorkflowNeedsATrigger
	}
	return action, nil
}

func (i interactor) findLinksForAction(workflowData domain.WorkflowContent, action *domain.WorkflowAction) []domain.BehaviourTreeLink {
	links := workflowData.Links
	outputLinks := []domain.BehaviourTreeLink{}

	for _, link := range links {
		if link.OutputAction.String() == action.ActionId {
			destinationAction, err := i.serviceGateway.GetService(link.InputAction)

			if err != nil {

			}

			outputLinks = append(outputLinks, domain.BehaviourTreeLink{
				DestinationAction: link.OutputAction,
				Link:              link,
			})
		}
	}
}

func (i interactor) buildBehaviourTree(workflowData domain.WorkflowContent) (*domain.BehaviourTreeAction, error) {
	trigger, err := i.findTriggerInWorkflow(workflowData)

	if err != nil {
		return nil, err
	}

	treeTrigger := domain.BehaviourTreeAction{
		Action: trigger,
		Links:  nil,
	}

}

func (i interactor) SetWorkflowData(user *domain.User, workflowId string, workflowData domain.WorkflowContent) error {
	workflow, err := i.workflowRepo.FindById(user.UserID, workflowId)

	if err != nil || workflow == nil {
		return domain.ErrWorkflowNotFound
	}

	i.buildBehaviourTree(workflowData)

	if err != nil {
		return err
	}

	return nil
}
