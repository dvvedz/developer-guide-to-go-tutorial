package main

import "fmt"

// interfaces makes it easier to reuse code
// We will build a englishBot and one swedishBot and resuse the code

type englishBot struct{}
type swedishBot struct{}

func main() {
	eb := englishBot{}
	sn := swedishBot{}
}

func printGreeting(eb englishBot) {
	fmt.Println(eb)
}

func printGreeting(sb swedishBot) {
	fmt.Println(sb)
}

func (eb englishBot) getGreeting() string {
	return "Hello World"
}

func (sb swedishBot) getGreeting() string {
	return "Hej Världen"
}
