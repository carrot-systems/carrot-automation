package domain

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

const (
	RUNNING_STATUS_STARTED = 0
	RUNNING_STATUS_DONE    = 1
	RUNNING_STATUS_ERRORED = 2
)

type RunningHistory struct {
	ID      uuid.UUID
	Started time.Time
	Done    sql.NullTime
	Status  int
	Error   string
}

func (runnigHistory RunningHistory) Initialize() {
	runnigHistory.ID = uuid.New()
}

type Link struct {
	Source         Action
	SourceOutputId string

	Destination        Action
	DestinationInputId string
}

type Workflow struct {
	ID      uuid.UUID
	Name    string
	User    *uuid.UUID
	History []*uuid.UUID
	Trigger *uuid.UUID
}

func (w *Workflow) Initialize() {
	w.ID = uuid.New()
}
