package domain

import "errors"

var (
	ErrServiceNotFound                     = errors.New("service not found")
	ErrFailedToGetUser                     = errors.New("failed to fetch user")
	ErrWorkflowAlreadyExistingWithThisName = errors.New("a workflow already exists with this name")
	ErrWorkflowNotFound                    = errors.New("workflow not found")
	ErrWorkflowNeedsATrigger               = errors.New("workflow needs a trigger")
	ErrWorkflowHasMoreThanOneTrigger       = errors.New("workflow has more than one trigger")
	ErrFeatureNotImplemented               = errors.New("feature not implemented")
)
