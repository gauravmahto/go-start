package utils

// Action constants
const (
	ActionEnter  = "enter"
	ActionExit   = "exit"
	ActionBefore = "before"
	ActionAfter  = "after"
)

// GetActionMsg Get message based on the action.
func GetActionMsg(action, name string) (message string) {

	switch action {

	case ActionEnter:
		message = "Enter #" + name

	case ActionExit:
		message = "Exit #" + name

	case ActionBefore:
		message = "Before #" + name

	case ActionAfter:
		message = "After #" + name

	default:
		message = "-----------------"

	}

	return

}
