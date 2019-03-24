package main

import "fmt"

const englishGreeting = "Hello, "
const germanGreeting = "Hallo, "
const frenchGreeting = "Bonjour, "
const spanishGreeting = "Hola, "

// Hello returns a personalized greeting
func Hello(name string, language string) string {
	// If no name, send greetings to the world
	if name == "" {
		name = "world"
	}

	// Greets in German
	if language == "German" {
		return germanGreeting + name
	}

	// Greets in French
	if language == "French" {
		return frenchGreeting + name
	}

	// Greets in Spanish
	if language == "Spanish" {
		return spanishGreeting + name
	}

	// Default
	return englishGreeting + name
}

func main() {
	fmt.Println(Hello("", ""))
}
