package main

import "./print"

const (
	Enter  = "enter"
	Exit   = "exit"
	Before = "before"
	After  = "after"
)

type salutation struct {
	name     string
	greeting string
}

func updateMessage(message *salutation) {

	// message.name = "Gaurav"
	// message.greeting = "Hello!"

	*message = salutation{name: "Gaurav", greeting: "Hello!"}

	// fmt.Println(*message, a, b)

}

func printMessage(message *salutation, printer print.Printer) {

	printer(message.greeting, message.name)

}

func getActionMsg(action, name string) (message string) {

	switch action {

	case Enter:
		message = "Enter #" + name

	case Exit:
		message = "Exit #" + name

	case Before:
		message = "Before #" + name

	case After:
		message = "After #" + name

	default:
		message = "-----------------"

	}

	return

}

func main() {

	var message = salutation{"User", "Hello GO"}
	// a, b := 1, 2 // Declaration and assignment.

	updateMsg := true

	if msg := getActionMsg(Before, "updateMessage"); updateMsg {

		print.PrintLine(msg)

	}

	printMessage(&message, print.PrintLine)

	if msg := getActionMsg(After, "updateMessage"); updateMsg {

		print.PrintLine(msg)

		updateMessage(&message)

	}

	printMessage(&message, print.PrintLine)

}
