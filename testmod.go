package testmod

import "fmt"
import "errors"

// Hi ...
func Hi(name string, lang string) (string, error) {
	switch lang {
	case "en":
		return fmt.Sprintf("Hello, %s", name), nil
	case "ru":
		return fmt.Sprintf("Привет, %s", name), nil
	default:
		return "", errors.New("Unknown language")
	}
}
