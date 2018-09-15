package greeting

import (
	"gmahto.local/utils"
)

// Message type
const (
	Informal = "informal"
	Formal   = "formal"
)

// CreateMessages Create a message.
func CreateMessages(name string) (message, formalMsg Salutation) {

	// var greetings map[string]string
	// greetings = make(map[string]string)
	// greetings[informal] = "Hi!"
	// greetings[formal] = "Hello!"

	greetings := map[string]string{
		Informal: "Hi!",
		Formal:   "Hello!",
	}

	message = Salutation{name: name, greeting: greetings[Informal]}

	formalMsg = Salutation{name, greetings[Formal]}

	return

}

// UpdateMessage Update the message.
func UpdateMessage(message *Salutation) {

	// *message = Salutation{name: "Gaurav", greeting: "Bye!"}

	message.greeting = "Bye!"

	// fmt.Println(*message, a, b)

}

// PrintMessage Prints a message.
func PrintMessage(message *Salutation, printer utils.Printer) {

	printer(message.greeting, message.name)

}

// PrintMessages Prints all the messages.
func PrintMessages(messages []*Salutation, printer utils.Printer) {

	for _, message := range messages {

		printer(message.greeting, message.name)

	}

}

// PrintMessageNTimes Prints a message N times.
func PrintMessageNTimes(message *Salutation, printer utils.Printer, times int) {

	for i := 0; i < times; i++ {

		printer(message.greeting, message.name)

	}

}
