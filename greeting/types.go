package greeting

import (
	"github.com/gauravmahto/go-start/utils"
)

// Salutation A greeting type.
type Salutation struct {
	name     string
	greeting string
}

// Message A Message interface.
type Message interface {
	UpdateMessage(greeting string)
	PrintMessage(printer utils.Printer)
	PrintMessageNTimes(printer utils.Printer, times int)
}

// Messages Collection of Message.
type Messages []Message
