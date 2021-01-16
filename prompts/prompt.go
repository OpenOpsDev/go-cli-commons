package prompts

import (
	"github.com/openopsdev/go-cli-commons/logger"
)

// CLIInput -
type CLIInput interface {
	Run() (string, error)
}

// Prompts - map of Promptui components
type Prompts map[string]CLIInput

// Answers -
type Answers map[string]string

// Run - starts the run
func (p Prompts) Run() Answers {
	answers := Answers{}

	for key, input := range p {
		a, err := input.Run()
		if err != nil {
			logger.Fatal(err.Error())
		}
		answers[key] = a
	}

	return answers
}
