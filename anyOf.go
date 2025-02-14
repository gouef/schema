package schema

import "errors"

// AnyOfField for anyOf
type AnyOfField struct {
	options      []Field
	defaultValue any
}

func (a *AnyOfField) Validate(value any) error {
	for _, t := range a.options {
		if err := t.Validate(value); err == nil {
			return nil
		}
	}
	return errors.New("value does not match any of the allowed types")
}

func (a *AnyOfField) Default(value any) Field {
	for _, opt := range a.options {
		if err := opt.Validate(value); err == nil {
			a.defaultValue = value
		}
	}
	return a
}

func (a *AnyOfField) HasDefault() bool {
	return a.defaultValue != nil
}

func (a *AnyOfField) GetDefault() any {
	return a.defaultValue
}

func (a *AnyOfField) CastTo(target any) (any, error) {
	return target, nil
}

func AnyOf(types ...Field) Field {
	return &AnyOfField{options: types}
}
