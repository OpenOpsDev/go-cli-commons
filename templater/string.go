package templater

import (
	"errors"

	"github.com/aymerick/raymond"
)

// StringTemplate -
type StringTemplate struct {
	TPL string
	Result string
}

// NewStringTemplate -
func NewStringTemplate(tpl string) *StringTemplate {
	return &StringTemplate {
		TPL: tpl,
	}
}

// Render -
func (s *StringTemplate) Render(ctx map[string]string) error {
	result, err := raymond.Render(s.TPL, ctx)

	if err != nil {
		return err
	}

	s.Result = result
	return nil
}

// Save - saves results to a file
func (s *StringTemplate) Save(dest string) error {
	if len(s.Result) > 0 {
		err := writeToFile(dest, s.Result)

		if err != nil {
			return err
		}
	} else {
		return errors.New("no result")
	}
	return nil
}