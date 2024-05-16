package errors

type IValidateError interface {
	Error() string
	Field() string
}

type ValidateError struct {
	Key string
	Err string
}

func (e ValidateError) Error() string {
	return e.Err
}

func (e ValidateError) Field() string {
	return e.Key
}
