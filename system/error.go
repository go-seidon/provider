package system

type Error struct {
	Code    int32
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(c int32, m string) *Error {
	return &Error{c, m}
}
