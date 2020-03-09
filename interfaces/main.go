package main

import "fmt"

//declaring an interface and abstract functions
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

//can pass the interface here
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//if you are implementing getGreeting method and same return type string it means you are implementing the interface
func (eb englishBot) getGreeting() string {
	return "Hello"
}

//same as in above comments
func (sb spanishBot) getGreeting() string {
	return "Hola"
}
