package exception

type ValidationError struct {
	Message string
}

func ThrowValidationError(message string) *error {

	var err error = &ValidationError{
		Message: message,
	}
	return &err
}

func (e *ValidationError) Error() string {
	return e.Message
}
