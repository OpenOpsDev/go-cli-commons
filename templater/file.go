package templater

import (
	"errors"

	"github.com/aymerick/raymond"
)

// FileMaker - Renders a template with passed vars
type FileMaker struct {
	Path string
	Result string
}

// NewFile -
func NewFile(path string) *FileMaker {
	return &FileMaker{
		Path: path,
	}
}

// Render -
func (f *FileMaker) Render(vars map[string]string) error {
	tpl, err := raymond.ParseFile(f.Path)
	if err != nil {
		return err
	}

	result, err := tpl.Exec(vars)

	if err != nil {
		return err
	}

	f.Result = result
	return nil
}

// Save - saves results to a file
func (f *FileMaker) Save(dest string) error {
	if len(f.Result) > 0 {
		err := writeToFile(dest, f.Result)

		if err != nil {
			return err
		}
	} else {
		return errors.New("no result")
	}
	return nil
}