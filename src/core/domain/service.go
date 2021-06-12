package domain

type Service struct {
	Id      string
	Name    string
	Version string
	Actions []Action `json:"-"` //excluding it from the json to not send massive responses, looking to the service directly will expose it
}

func (s Service) GetActions() []Action {
	return s.Actions
}
