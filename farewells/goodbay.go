package farewells

import "fmt"
import "errors"

// Goodbay ...
func Goodbay(name string, lang string) (string, error) {
	switch lang {
	case "en":
		return fmt.Sprintf("Goodbay, %s!", name), nil
	case "ru":
		return fmt.Sprintf("Пока, %s!", name), nil
	default:
		return "", errors.New("Unknown language")
	}
}
