package greetings

import (
	"errors"
	"fmt"
)

// Returns a greeting for the named person
func Hello(name string) (string, error) {
	// if no name was given, return an error
	if name == "" {
		return "", errors.New("Empty name")
	}
	message := fmt.Sprintf("Hello, %v. Welcome", name)
	return message, nil
}
