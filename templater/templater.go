package templater

import "os"

// Templater -
type Templater interface {
	Render(vars map[string]string) error
	Save(dest string) error
}

func writeToFile(dest, result string) error {
	file, err := os.Create(dest)

    if err != nil {
        return err
    }

    defer file.Close()

    _, err = file.WriteString(result)

    if err != nil {
        return err
	}
	
	return nil
}