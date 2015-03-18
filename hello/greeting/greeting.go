package greeting

import "fmt"

// Salutation is an exported type.
type Salutation struct {
	Name     string
	Greeting string
}

// Renamable is an exported type.
type Renamable interface {
	Rename(newName string)
}

// Rename is an exported function.
func (salutation *Salutation) Rename(newName string) {
	salutation.Name = newName
}

func (salutation *Salutation) Write(p []byte) (n int, err error) {
	s := string(p)
	salutation.Rename(s)
	n = len(s)
	err = nil
	return
}

// Salutations is an exported type.
type Salutations []Salutation

// Printer is an exported type.
type Printer func(string)

// Greet is an exported function.
func (salutations Salutations) Greet(do Printer, isFormal bool, times int) {
	for _, s := range salutations {
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
	prefixMap := map[string]string{
		"Bob":  "Mr. ",
		"Joe":  "Dr. ",
		"Amy":  "Dr. ",
		"Mary": "Mrs. ",
	}

	prefixMap["Joe"] = "Jr "

	if _, exists := prefixMap["Mary"]; exists {
		delete(prefixMap, "Mary")
	}

	if value, exists := prefixMap[name]; exists {
		return value
	}

	return "Dude "
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
