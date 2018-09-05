package print

import "fmt"

// Printer Printer type.
type Printer func(...string)

// Print Prints a message.
func Print(message ...string) {

	fmt.Print(message)

}

// PrintLine Prints a message on a new line.
func PrintLine(message ...string) {

	fmt.Println(message)

}
