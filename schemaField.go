package schema

import (
	"reflect"
	"strings"
)

type Field interface {
	Validate(value any) error
}

// Normalize function to clean and prepare data
func Normalize(data any) (any, error) {
	v := reflect.ValueOf(data)

	switch v.Kind() {
	case reflect.String:
		return strings.TrimSpace(v.String()), nil
	case reflect.Int:
		return v.Int(), nil
	case reflect.Bool:
		return v.Bool(), nil
	case reflect.Float64:
		return v.Float(), nil
	}

	return data, nil
}

// Process function now uses Normalize
func Process(schema Field, data any) error {
	normalizedData, err := Normalize(data)
	if err != nil {
		return err
	}

	return schema.Validate(normalizedData)
}
