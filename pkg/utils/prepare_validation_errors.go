package utils

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

func PrepareValidationErrors(errors validator.ValidationErrors, requestStruct interface{}) map[string]string {
	errorMessages := make(map[string]string)
	for _, fieldErr := range errors {
		structField, _ := reflect.TypeOf(requestStruct).FieldByName(fieldErr.StructField())
		jsonFieldName := structField.Tag.Get("json")
		message := fieldErr.Field() + " " + fieldErr.Tag() + " " + fieldErr.Param()
		errorMessages[jsonFieldName] = message
	}
	return errorMessages
}
