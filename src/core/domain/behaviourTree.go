package domain

type BehaviourTreeAction struct {
	Action *WorkflowAction
	Links  []*BehaviourTreeLink
}

type BehaviourTreeLink struct {
	DestinationAction *BehaviourTreeAction
	Link              map[string]string //Source output to destination input
}
