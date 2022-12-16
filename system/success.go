package system

type Success struct {
	Code    int32
	Message string
}

func NewSuccess(c int32, m string) Success {
	return Success{c, m}
}
