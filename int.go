package schema

import "errors"

type IntField struct{}

func (i IntField) Validate(value any) error {
	if _, ok := value.(int); !ok {
		return errors.New("expected integer")
	}
	return nil
}

func Int() Field {
	return IntField{}
}
