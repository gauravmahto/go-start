package greeting

import (
	"github.com/gauravmahto/go-start/utils"
)

// Message tone
const (
	Informal = "informal"
	Formal   = "formal"
)

// CreateMessages Create a message.
func CreateMessages(name string) (message, formalMsg Message) {

	// var greetings map[string]string
	// greetings = make(map[string]string)
	// greetings[informal] = "Hi!"
	// greetings[formal] = "Hello!"

	greetings := map[string]string{
		Informal: "Hi!",
		Formal:   "Hello!",
	}

	message = &Salutation{name: name, greeting: greetings[Informal]}

	formalMsg = &Salutation{name, greetings[Formal]}

	return

}

// PrintMessages Prints all the messages.
func PrintMessages(messages Messages, printer utils.Printer) {

	for _, message := range messages {

		message.PrintMessage(printer)

	}

}

// UpdateMessage Update the message.
func (message *Salutation) UpdateMessage(greeting string) {

	// *message = Salutation{name: "Gaurav", greeting: "Bye!"}

	message.greeting = greeting

	// fmt.Println(*message, a, b)

}

// PrintMessage Prints a message.
func (message *Salutation) PrintMessage(printer utils.Printer) {

	printer(message.greeting, message.name)

}

// PrintMessageNTimes Prints a message N times.
func (message *Salutation) PrintMessageNTimes(printer utils.Printer, times int) {

	for i := 0; i < times; i++ {

		printer(message.greeting, message.name)

	}

}
