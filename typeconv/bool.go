package typeconv

func Bool(i bool) *bool {
	return &i
}

func BoolVal(i *bool) bool {
	if i == nil {
		return false
	}
	return *i
}
