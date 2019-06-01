package main

import (
	"github.com/gauravmahto/go-start/greeting"
	"github.com/gauravmahto/go-start/utils"
)

// Entry point.
func main() {

	var message, formalMsg = greeting.CreateMessages("Gaurav")

	updateMsg := true

	if msg := utils.GetActionMsg(utils.ActionBefore, "updateMessage"); updateMsg {

		utils.PrintLine(msg)

	}

	message.PrintMessage(utils.PrintLine)

	if msg := utils.GetActionMsg(utils.ActionAfter, "updateMessage"); updateMsg {

		utils.PrintLine(msg)

		message.UpdateMessage("Bye!")

	}

	message.PrintMessage(utils.PrintLine)

	// A non-bufferred channel named 'done' and accepting a boolean value.
	done := make(chan bool)

	// A go routine.
	go func() { // Anonymous function, acting as a closure.

		utils.PrintLine("-- GO ROUTINE --")
		message.PrintMessageNTimes(utils.PrintLine, 1)
		utils.PrintLine("-- GO ROUTINE --")

		done <- true // Send boolean true to the 'done' channel.

	}() // Execute the anonymous function.

	messages := greeting.Messages{message, formalMsg}
	greeting.PrintMessages(messages, utils.PrintLine)

	// Read from 'done' channel. Will block the current execution,
	// until data is available on the 'done' channel to be read.
	<-done

}
