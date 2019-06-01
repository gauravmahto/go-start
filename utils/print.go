package utils

import (
	"fmt"
)

// Printer Printer type.
type Printer func(...interface{})

// Print Prints a message.
func Print(message ...interface{}) {

	fmt.Print(message...)

}

// PrintLine Prints a message on a new line.
func PrintLine(message ...interface{}) {

	fmt.Println(message...)

}
