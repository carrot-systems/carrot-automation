package standardtools

import "github.com/carrot-systems/carrot-automation/src/core/domain"

func NewStandardTools() domain.Service {
	return domain.Service{
		Id:      "fr.adinunno.carrotautomation.standardtool",
		Name:    "Standard tools",
		Version: "0.1",
		Actions: []domain.Action{
			{
				ActionId:    "five",
				IsTrigger:   true,
				Name:        "5 minutes timer",
				Description: "Executes the workflow every 5 minutes",
				Input:       nil,
				Output: []domain.IO{
					{
						IOId:        "output",
						Name:        "Output",
						Description: "",
						Type:        "void",
					},
				},
			}, {
				ActionId:    "three",
				IsTrigger:   true,
				Name:        "3 minutes timer",
				Description: "Executes the workflow every 3 minutes",
				Input:       nil,
				Output: []domain.IO{
					{
						IOId:        "output",
						Name:        "Output",
						Description: "",
						Type:        "void",
					},
				},
			}, {
				ActionId:    "one",
				IsTrigger:   true,
				Name:        "1 minutes timer",
				Description: "Executes the workflow every minute",
				Input:       nil,
				Output: []domain.IO{
					{
						IOId:        "output",
						Name:        "Output",
						Description: "",
						Type:        "void",
					},
				},
			},
		},
	}
}
