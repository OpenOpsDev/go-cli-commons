package prompts

import (
	"errors"
	"regexp"
	"strings"

	"github.com/manifoldco/promptui"
)

type MultiInput struct {
	Label       string
	ValidString string
}

func (m *MultiInput) Run() (string, error) {
	keepGoing := true
	var answers []string

	prompt := promptui.Prompt{
		Label: m.Label,
		Validate: func(input string) error {
			if len(m.ValidString) > 0 {
				if matched, err := regexp.Match(m.ValidString, []byte(input)); err != nil || !matched {
					return errors.New("not a valid input")
				}
			}

			return nil
		},
	}

	for keepGoing {
		result, err := prompt.Run()

		if err != nil {
			return "", err
		}

		answers = append(answers, result)

		add := promptui.Select{
			Label: "Add more? [Y",
			Items: []string{"yes", "no"},
		}
		_, confirm, err := add.Run()
		if err != nil {
			return "", err
		}

		if confirm == "no" {
			keepGoing = false
		}
	}

	return strings.Join(answers, ","), nil
}
