package main

import "github.com/hdvegasilva/greeting"

func main() {

	var s = greeting.Salutation{"Bob", "Hello"}
	greeting.Greet(s, greeting.CreatePrintCustom("?"), true, 5)
}
