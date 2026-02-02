package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	//Set properties of the predefines logger, including
	//the log entry prefix and a flag to disable printing
	// the time , source file and the line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.

	names := []string{
		"Gladys",
		"Samantha",
		"Darrin"}
	//Request a greeting message.
	messages, err := greetings.Hellos(names)
	//if an error was returned, print it to the console and
	//exit the program.
	if err != nil {
		log.Fatal(err)
	}

	//else if no error , print the return message to the console.
	fmt.Println(messages)
}
