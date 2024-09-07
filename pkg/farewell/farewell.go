package farewell

import (
	"errors"
	"fmt"
	"math/rand"
)

// SayHello returns a greeting for the given name.
func One(name string) (string, error) {
	if name == "" {
		return "", errors.New("parameter 'name' is empty")
	}

	if name == "error" {
		return "", fmt.Errorf("error: %v", name)
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Until next time, %v. Bye!",
		"Take care, %v!",
		"Farewell, %v! It was my pleasure!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}

// Many returns a map that associates each of the named people
func Many(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := One(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}
