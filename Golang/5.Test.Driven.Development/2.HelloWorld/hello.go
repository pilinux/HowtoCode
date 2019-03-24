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

	// Default greeting
	greeting := englishGreeting

	switch language {
	// Greets in German
	case "German":
		greeting = germanGreeting
	// Greets in French
	case "French":
		greeting = frenchGreeting
	// Greets in Spanish
	case "Spanish":
		greeting = spanishGreeting
	}

	return greeting + name
}

func main() {
	fmt.Println(Hello("", ""))
}
