package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/wendryuslima/user-manager/src/configuration/rest_err"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		enLocale := en.New()
		universalTranslator := ut.New(enLocale, enLocale)
		transl, _ = universalTranslator.GetTranslator("en")
		_ = en_translations.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validationErr error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var validationErrors validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
	}

	if errors.As(validationErr, &validationErrors) {
		errorCauses := make([]rest_err.Causes, 0, len(validationErrors))

		for _, validationError := range validationErrors {
			errorCauses = append(errorCauses, rest_err.Causes{
				Message: validationError.Translate(transl),
				Field:   validationError.Field(),
			})
		}

		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorCauses)
	}

	return rest_err.NewBadRequestError("Error trying to convert fields")
}
