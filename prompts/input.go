package prompts

import "github.com/manifoldco/promptui"

type UserInput struct {
	Label string
}

func (u *UserInput) Run() (string, error) {
	prompt := promptui.Prompt{
		Label:    u.Label,
	}

	return prompt.Run()
}