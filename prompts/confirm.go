package prompts

import "github.com/manifoldco/promptui"

// ConfirmInput -
type ConfirmInput struct {
	Label string
}

// Run -
func (c ConfirmInput) Run() (string, error) {
	prompt := promptui.Prompt{
		Label:     c.Label,
		IsConfirm: true,
	}

	return prompt.Run()
}