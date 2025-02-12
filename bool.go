package schema

import "errors"

type BoolField struct{}

func (b BoolField) Validate(value any) error {
	if _, ok := value.(bool); !ok {
		return errors.New("expected boolean")
	}
	return nil
}

func Bool() Field {
	return BoolField{}
}
