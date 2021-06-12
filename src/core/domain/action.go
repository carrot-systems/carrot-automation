package domain

import "github.com/google/uuid"

type IO struct {
	ID uuid.UUID

	IOId        string
	Name        string
	Description string
	Type        string
}

func (io IO) Initialize() {
	io.ID = uuid.New()
}

type Action struct {
	ID uuid.UUID

	ActionId    string
	IsTrigger   bool
	Name        string
	Description string
	Input       []IO
	Output      []IO
}

func (action Action) Initialize() {
	action.ID = uuid.New()
}
