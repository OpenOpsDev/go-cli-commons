package prompts

import "github.com/manifoldco/promptui"

type SelectInput struct {
	Label   string
	Options []string
}

func (s *SelectInput) Run() (string, error) {
	prompt := promptui.Select{
		Label: s.Label,
		Items: s.Options,
	}

	_, a, err := prompt.Run()
	return a, err
}
