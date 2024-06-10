package main

import "fmt"

const (
	spanish = "Spanish"
	french = "French"
	german = "German"

	spanishGreeting = "Hola, "
	frenchGreeting = "Salut, "
	germanGreeting = "Guten Tag, "
	defaultGreeting = "Hello, "

	defaultName = "World"
)

func Hello(name string, language string) string {
	if name == ""{
		name = defaultName
	}

	return greeting(language)+name
}

func greeting(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = spanishGreeting
	case "French":
		prefix = frenchGreeting
	case "German":
		prefix = germanGreeting
	default:
		prefix = defaultGreeting
	}
	return
}

func main() {
	fmt.Println(Hello("Charlie", ""))
}