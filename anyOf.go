package schema

import "errors"

// AnyOfField for anyOf
type AnyOfField struct {
	types []Field
}

func (a AnyOfField) Validate(value any) error {
	for _, t := range a.types {
		if err := t.Validate(value); err == nil {
			return nil
		}
	}
	return errors.New("value does not match any of the allowed types")
}

func AnyOf(types ...Field) Field {
	return AnyOfField{types: types}
}
