package system

type SystemError struct {
	Code    int32
	Message string
}

func (e *SystemError) Error() string {
	return e.Message
}

func NewError(c int32, m string) *SystemError {
	return &SystemError{c, m}
}
