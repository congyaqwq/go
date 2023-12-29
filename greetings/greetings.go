package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name cannot be empty")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{"Hello, %s!", "Hi, %s!", "Greetings, %s!"}

	return formats[rand.Intn(len(formats))]
}
