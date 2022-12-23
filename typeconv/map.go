package typeconv

import (
	"reflect"
)

func MapVal(i map[string]interface{}) map[string]interface{} {
	r := map[string]interface{}{}
	if i == nil {
		return r
	}
	for k, v := range i {
		r[k] = v
	}

	for k, v := range r {
		isPtrNil := reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil()
		if v == nil || isPtrNil {
			delete(r, k)
		}
	}
	return r
}
