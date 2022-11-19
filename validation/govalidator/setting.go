package govalidator

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type CustomValidator = func(fl validator.FieldLevel) bool
type CustomTagName = func(field reflect.StructField) string

func FieldTagName(n string) CustomTagName {
	return func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get(n), ",", 2)[0]
		if name == "-" || name == "" {
			return field.Name
		}
		return name
	}
}
