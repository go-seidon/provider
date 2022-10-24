package error

type SystemError struct {
	Code    int32
	Message string
}

func (e *SystemError) Error() string {
	return e.Message
}

func New(c int32, m string) *SystemError {
	return &SystemError{c, m}
}
