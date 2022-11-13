package typeconv

func Int(i int) *int {
	return &i
}

func IntVal(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

func Int32(i int32) *int32 {
	return &i
}

func Int32Val(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}

func Int64(i int64) *int64 {
	return &i
}

func Int64Val(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}
