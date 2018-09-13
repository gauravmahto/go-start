package main

import "gmahto.local/print"

// Action constants
const (
	Enter  = "enter"
	Exit   = "exit"
	Before = "before"
	After  = "after"
)

const (
	informal = "informal"
	formal   = "formal"
)

func createMessages(name string) (message, formalMsg Salutation) {

	// var greetings map[string]string
	// greetings = make(map[string]string)
	// greetings[informal] = "Hi!"
	// greetings[formal] = "Hello!"

	greetings := map[string]string{
		informal: "Hi!",
		formal:   "Hello!",
	}

	message = Salutation{name: name, greeting: greetings[informal]}

	formalMsg = Salutation{name, greetings[formal]}

	return

}

func updateMessage(message *Salutation) {

	// *message = Salutation{name: "Gaurav", greeting: "Bye!"}

	message.greeting = "Bye!"

	// fmt.Println(*message, a, b)

}

func printMessage(message *Salutation, printer print.Printer) {

	printer(message.greeting, message.name)

}

func printMessages(messages []*Salutation, printer print.Printer) {

	for _, message := range messages {

		printer(message.greeting, message.name)

	}

}

func printMessageNTimes(message *Salutation, printer print.Printer, times int) {

	for i := 0; i < times; i++ {

		printer(message.greeting, message.name)

	}

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

	var message, formalMsg = createMessages("Gaurav")
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

	// printMessageNTimes(&message, print.PrintLine, 1)

	printMessages([]*Salutation{&message, &formalMsg}, print.PrintLine)

}
