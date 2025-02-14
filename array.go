package schema

import (
	"errors"
	"reflect"
)

// ArrayField for handling slices
type ArrayField struct {
	elem         Field
	defaultValue []any
}

func (a *ArrayField) Validate(value any) error {
	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Map {
		return errors.New("expected array or map")
	}
	if v.Kind() == reflect.Map {
		// If it's a map, we validate it like an array of key-value pairs
		for _, key := range v.MapKeys() {
			if err := a.elem.Validate(v.MapIndex(key).Interface()); err != nil {
				return errors.New("map value " + err.Error())
			}
		}
	} else {
		// Validate elements in the array
		for i := 0; i < v.Len(); i++ {
			if err := a.elem.Validate(v.Index(i).Interface()); err != nil {
				return errors.New("array element " + string(rune(i)) + ": " + err.Error())
			}
		}
	}
	return nil
}

func (a *ArrayField) Default(value any) Field {
	if v, ok := value.([]any); ok {
		a.defaultValue = v
	}
	return a
}

func (a *ArrayField) CastTo(target any) (any, error) {
	return target, nil
}

func (a *ArrayField) HasDefault() bool {
	return len(a.defaultValue) > 0
}

func (a *ArrayField) GetDefault() any {
	return a.defaultValue
}

func ArrayOf(elem Field) Field {
	return &ArrayField{elem: elem}
}

func Array(elem Field) Field {
	return &ArrayField{elem: elem}
}

// Helper function to check if a value matches any of the provided types
func validateAnyOf(value any, variants ...any) error {
	for _, v := range variants {
		if reflect.DeepEqual(value, v) {
			return nil
		}
	}
	return errors.New("value does not match any of the allowed types")
}
