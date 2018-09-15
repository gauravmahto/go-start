package main

import (
	"gmahto.local/greeting"
	"gmahto.local/utils"
)

// Action constants
const (
	Enter  = "enter"
	Exit   = "exit"
	Before = "before"
	After  = "after"
)

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

	var message, formalMsg = greeting.CreateMessages("Gaurav")
	// a, b := 1, 2 // Declaration and assignment.

	updateMsg := true

	if msg := getActionMsg(Before, "updateMessage"); updateMsg {

		utils.Line(msg)

	}

	greeting.PrintMessage(&message, utils.Line)

	if msg := getActionMsg(After, "updateMessage"); updateMsg {

		utils.Line(msg)

		greeting.UpdateMessage(&message)

	}

	greeting.PrintMessage(&message, utils.Line)

	// greeting.PrintMessageNTimes(&message, utils.Line, 1)

	greeting.PrintMessages([]*greeting.Salutation{&message, &formalMsg}, utils.Line)

}
