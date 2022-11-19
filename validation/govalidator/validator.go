package govalidator

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/go-seidon/provider/validation"
)

type goValidator struct {
	client     *validator.Validate
	translator ut.Translator
}

// @note: returning first invalid error
func (v *goValidator) Validate(i interface{}) error {
	err := v.client.Struct(i)
	if err == nil {
		return nil
	}

	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return &validation.ValidationError{
			Message: err.Error(),
		}
	}

	if len(errs) == 0 {
		return &validation.ValidationError{
			Message: err.Error(),
		}
	}

	return &validation.ValidationError{
		Message: errs[0].Translate(v.translator),
	}
}

func NewValidator() *goValidator {
	en := en.New()
	uniTrans := ut.New(en, en)
	translator, _ := uniTrans.GetTranslator("en")

	client := validator.New()

	en_translations.RegisterDefaultTranslations(client, translator)
	client.RegisterTagNameFunc(FieldTagName("label"))

	return &goValidator{
		client:     client,
		translator: translator,
	}
}
