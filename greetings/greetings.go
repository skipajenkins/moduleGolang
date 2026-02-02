package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

//Hello returns a greeting for the named person.

func Hello(name string) (string, error) {
	//if no name was given, return an error with a messaege.
	if name == "" {
		return "", errors.New("empty name")
	}
	//(Basically we created an error handler above which tells us
	// how to respond when the name parameter is empty.)

	// if a name was received, return a value that embeds the name
	// in a greeting message.

	// Create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil // Return the message with the name or return nil

}

// Hellos returns a map that associates each of the named people
// with a gretting message.
func Hellos(name []string) (map[string]string, error) {
	//A map to associate anmes with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range name {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map. associate the retrueved message with
		// the name.
		messages[name] = message
	}
	return messages, nil
}

//randomFormat reutnrns one of a set of greeting messages, The returned
//message is selected at random.

func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	// Return a randomly selected message format by specifying
	// a random index for the slice of formats
	return formats[rand.Intn(len(formats))]
}
