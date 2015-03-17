package greeting

import "fmt"

// Salutation is an exported type.
type Salutation struct {
	Name     string
	Greeting string
}

// Printer is an exported type.
type Printer func(string)

// Greet is an exported function.
func Greet(salutation []Salutation, do Printer, isFormal bool, times int) {
	for _, s := range salutation {
		message, alternate := CreateMessage(s.Name, s.Greeting)
		if prefix := GetPrefix(s.Name); isFormal {
			do(prefix + message)
		} else {
			do(alternate)
		}
	}
}

// GetPrefix is an exported function.
func GetPrefix(name string) (prefix string) {
	switch {
	case name == "Bob":
		prefix = "Mr. "
	case name == "Joe", name == "Amy", len(name) == 10:
		prefix = "Dr. "
	case name == "Mary":
		prefix = "Mrs. "
	default:
		prefix = "Dude "
	}

	return
}

// TypeSwitchTest is an exported function.
func TypeSwitchTest(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Salutation:
		fmt.Println("salutation")
	default:
		fmt.Println("unknown")
	}
}

// CreateMessage is an exported function.
func CreateMessage(name string, greeting string) (message string, alternate string) {
	message = greeting + " " + name
	alternate = "HEY!" + " " + name
	return
}

// Print is an exported function.
func Print(s string) {
	fmt.Print(s)
}

// PrintLine is an exported function.
func PrintLine(s string) {
	fmt.Println(s)
}

// CreatePrintCustom is an exported function.
func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}
