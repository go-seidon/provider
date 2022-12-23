package typeconv

func MapVal(i map[string]interface{}) map[string]interface{} {
	if i == nil {
		return map[string]interface{}{}
	}
	for k, v := range i {
		if v == nil {
			delete(i, k)
		}
	}
	return i
}
