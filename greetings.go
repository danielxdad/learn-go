package main

import (
	"errors"
	"fmt"
	"math/rand"
	// "reflect"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	return fmt.Sprintf(randomFormat(), name), nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string, len(names))

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// fmt.Println(reflect.TypeOf(formats[len(formats)-1]))

	return formats[rand.Intn(len(formats))]
}
