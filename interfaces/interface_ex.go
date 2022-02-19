package interfaces

import "fmt"

type bot interface {
	getGreeting(string) string
}

type englishBot struct{}

type spanishBot struct{}

func (englishBot) getGreeting(name string) string {
	return "Hello " + name
}

func (spanishBot) getGreeting(name string) string {
	return name + " Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting("Malli"))
}
