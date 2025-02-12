package schema

import "errors"

type StringField struct{}

func (b StringField) Validate(value any) error {
	if _, ok := value.(string); !ok {
		return errors.New("expected string")
	}
	return nil
}

func String() Field {
	return StringField{}
}
