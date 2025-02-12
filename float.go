package schema

import "errors"

type FloatField struct{}

func (f FloatField) Validate(value any) error {
	if _, ok := value.(float64); !ok {
		return errors.New("expected float64")
	}
	return nil
}

func Float() Field {
	return FloatField{}
}
