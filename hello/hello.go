package main

import "github.com/hdvegasilva/greeting"

func main() {
	//var s = greeting.Salutation{"Bob", "Hello"}

	slice := []greeting.Salutation{
		{"Bob", "Hello"},
		{"Joe", "Hi"},
		{"Mary", "What is up?"},
	}

	greeting.Greet(slice, greeting.CreatePrintFunction("?"), true, 5)
}
