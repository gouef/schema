package schema

import "errors"

type BoolField struct {
	defaultValue bool
}

func (b *BoolField) Validate(value any) error {
	if _, ok := value.(bool); !ok {
		return errors.New("expected boolean")
	}
	return nil
}

func (b *BoolField) Default(value any) Field {
	if v, ok := value.(bool); ok {
		b.defaultValue = v
	}
	return b
}

func (b *BoolField) CastTo(target any) (any, error) {
	return target, nil
}

func Bool() Field {
	return &BoolField{}
}
