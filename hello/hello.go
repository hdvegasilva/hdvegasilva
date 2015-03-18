package main

import "./greeting"
import "fmt"

func RenameToFrog(r greeting.Renamable) {
	r.Rename("Frog")
}

func main() {
	//var s = greeting.Salutation{"Bob", "Hello"}

	salutations := greeting.Salutations{
		{"Bob", "Hello"},
		{"Joe", "Hi"},
		{"Mary", "What is up?"},
	}

	//salutations[0].Rename("John")

	//RenameToFrog(&salutations[0])

	fmt.Fprintf(&salutations[0], "The count is %d", 10)

	salutations.Greet(greeting.CreatePrintFunction("?"), true, 6)
	//greeting.Greet(slice, greeting.CreatePrintFunction("?"), true, 5)
}
