package typeconv

func Float32(i float32) *float32 {
	return &i
}

func Float32Val(i *float32) float32 {
	if i == nil {
		return 0
	}
	return *i
}

func Float64(i float64) *float64 {
	return &i
}

func Float64Val(i *float64) float64 {
	if i == nil {
		return 0
	}
	return *i
}
