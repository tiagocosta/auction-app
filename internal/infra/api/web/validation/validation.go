package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	validator_en "github.com/go-playground/validator/v10/translations/en"
	"github.com/tiagocosta/auction-app/configuration/rest_err"
)

var (
	Validate   = validator.New()
	translator ut.Translator
)

func init() {
	if value, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		enTranslation := ut.New(en, en)
		translator, _ = enTranslation.GetTranslator("en")
		validator_en.RegisterDefaultTranslations(value, translator)
	}
}

func ValidateErr(validation_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidation validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("invalid type error")
	} else if errors.As(validation_err, &jsonValidation) {
		causes := []rest_err.Cause{}
		for _, e := range validation_err.(validator.ValidationErrors) {
			causes = append(causes, rest_err.Cause{
				Field:   e.Field(),
				Message: e.Translate(translator),
			})
		}
		return rest_err.NewBadRequestError("invalid field values", causes...)
	} else {
		return rest_err.NewBadRequestError("error trying to convert fields")
	}
}
