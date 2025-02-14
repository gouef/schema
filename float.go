package schema

import "errors"

type FloatField struct {
	defaultValue float64
}

func (f *FloatField) Validate(value any) error {
	if _, ok := value.(float64); !ok {
		return errors.New("expected float64")
	}
	return nil
}

func (f *FloatField) Default(value any) Field {
	if v, ok := value.(float64); ok {
		f.defaultValue = v
	}
	return f
}

func (f *FloatField) CastTo(target any) (any, error) {
	return target, nil
}

func Float() Field {
	return &FloatField{}
}
