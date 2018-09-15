package print

import (
	"fmt"
)

// Printer Printer type.
type Printer func(...interface{})

// Print Prints a message.
func Print(message ...interface{}) {

	fmt.Print(message...)

}

// Line Prints a message on a new line.
func Line(message ...interface{}) {

	fmt.Println(message...)

}
