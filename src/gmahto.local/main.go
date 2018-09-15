package main

import (
	"gmahto.local/greeting"
	"gmahto.local/utils"
)

// Entry point.
func main() {

	var message, formalMsg = greeting.CreateMessages("Gaurav")

	updateMsg := true

	if msg := utils.GetActionMsg(utils.ActionBefore, "updateMessage"); updateMsg {

		utils.Line(msg)

	}

	message.PrintMessage(utils.Line)

	if msg := utils.GetActionMsg(utils.ActionAfter, "updateMessage"); updateMsg {

		utils.Line(msg)

		message.UpdateMessage("Bye!")

	}

	message.PrintMessage(utils.Line)

	// message.PrintMessageNTimes(utils.Line, 1)

	messages := greeting.Messages{message, formalMsg}
	greeting.PrintMessages(messages, utils.Line)

}
