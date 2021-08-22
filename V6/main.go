package main

import "fmt"

type Bot interface {
	GetGreeting(...interface{}) string
}

type EnglishBot struct{}
type SpanishBot struct{}
type PortugueseBot struct{}

func init() {}

func main() {
	eb := EnglishBot{}
	// sp := SpanishBot{}
	pb := PortugueseBot{}

	PrintGreeting(eb)
	PrintGreeting(pb)

	// This will fail because we're missing GetGreeting
	// Golang will look inside the current type looking for the GetGreeting function
	// TLDR; Anyone who implements GetGreeting can call the PrintGreeting function
	// PrintGreeting(sp)
}

func PrintGreeting(b Bot) {
	fmt.Println(b.GetGreeting())
}

func (EnglishBot) GetGreeting(...interface{}) string {
	return "Hello there!"
}

func (PortugueseBot) GetGreeting(values ...interface{}) string {
	for index, value := range values {
		fmt.Println(index, value)
	}
	return "Olá!"
}
