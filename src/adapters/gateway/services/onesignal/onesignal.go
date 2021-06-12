package onesignal

import "github.com/carrot-systems/carrot-automation/src/core/domain"

func NewOnesignal() domain.Service {
	return domain.Service{
		Id:      "fr.adinunno.carrotautomation.onesignal",
		Name:    "Onesignal",
		Version: "0.1",
		Actions: []domain.Action{
			{
				ActionId:    "notification",
				IsTrigger:   true,
				Name:        "Notification",
				Description: "Sends a notification",
				Output:      nil,
				Input: []domain.IO{
					{
						IOId:        "input",
						Name:        "input",
						Description: "",
						Type:        "void",
					},
				},
			},
		},
	}
}
