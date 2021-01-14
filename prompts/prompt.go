package prompts

import (
	"fmt"
	"os"
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
			fmt.Errorf("error with prompt: ", err)
			os.Exit(1)
		}
		answers[key] = a
	}

	return answers
}
