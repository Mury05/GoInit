package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Returns a greeting for the named person
func Hello(name string) (string, error) {
	// if no name was given, return an error
	if name == "" {
		return "", errors.New("Empty name")
	}
	message := fmt.Sprintf(RandomFormat(), name)
	return message, nil
}

// RandomFormat returns one of a set of greeting messages
func RandomFormat() string {
	// A slice of message format
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
