package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spainshBot struct{}

func main() {
	eb := englishBot{}
	sb := spainshBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	//Very custom logic for generating an english greeting
	return "Hi There!"
}

func (spainshBot) getGreeting() string {
	return "Hola!"
}
