package domain

import "fmt"

type ValidationError struct {
	Err string
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s", v.Err)
}
