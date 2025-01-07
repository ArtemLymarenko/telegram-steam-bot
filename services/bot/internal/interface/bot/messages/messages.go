package messages

import "fmt"

type Message string

const (
	MessageGreeting Message = "Welcome to out Steam Bot!"
)

func (m Message) String() string {
	return string(m)
}

func (m Message) Format(format string, msg string) string {
	return fmt.Sprintf(format, m.String(), msg)
}
