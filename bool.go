package schema

import "errors"

type BoolField struct {
	defaultValue bool
	required     bool
}

func (b *BoolField) Validate(value any) error {
	if _, ok := value.(bool); !ok {
		return errors.New("expected boolean")
	}
	return nil
}

func (b *BoolField) Required() Field {
	b.required = true
	return b
}

func (b *BoolField) IsRequired() bool {
	return b.required
}

func (b *BoolField) Default(value any) Field {
	if v, ok := value.(bool); ok {
		b.defaultValue = v
	}
	return b
}

func (b *BoolField) HasDefault() bool {
	return b.defaultValue != false
}

func (b *BoolField) GetDefault() any {
	return b.defaultValue
}

func (b *BoolField) CastTo(target any) (any, error) {
	return target, nil
}

func Bool() Field {
	return &BoolField{required: false}
}
