package domain

type ValidationError interface {
	Error() string
}
