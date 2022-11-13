package typeconv

func String(i string) *string {
	return &i
}

func StringVal(i *string) string {
	if i == nil {
		return ""
	}
	return *i
}
