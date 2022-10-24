package system

type SystemSuccess struct {
	Code    int32
	Message string
}

func NewSuccess(c int32, m string) *SystemSuccess {
	return &SystemSuccess{c, m}
}
