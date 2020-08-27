package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	printGreeting(englishBot{})
	printGreeting(spanishBot{})
}
