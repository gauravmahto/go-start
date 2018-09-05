package main

import "./print"

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

func main() {

	var message = salutation{"User", "Hello GO"}
	// a, b := 1, 2 // Declaration and assignment.

	print.PrintLine("Before calling #test()")

	printMessage(&message, print.PrintLine)

	print.PrintLine("After calling #test()")

	updateMessage(&message)

	printMessage(&message, print.PrintLine)

}
