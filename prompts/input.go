package prompts

import "github.com/manifoldco/promptui"

type UserInput struct {
	Label   string
	Default string
}

func (u *UserInput) Run() (string, error) {
	prompt := promptui.Prompt{
		Label: u.Label,
	}

	if len(u.Default) > 0 {
		prompt.Default = u.Default
	}

	return prompt.Run()
}
