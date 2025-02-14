package schema

import "errors"

type StringField struct {
	defaultValue string
}

func (b *StringField) Validate(value any) error {
	if _, ok := value.(string); !ok {
		return errors.New("expected string")
	}
	return nil
}

func (b *StringField) Default(value any) Field {
	if v, ok := value.(string); ok {
		b.defaultValue = v
	}
	return b
}

func (b *StringField) CastTo(target any) (any, error) {
	return target, nil
}

func String() Field {
	return &StringField{}
}
