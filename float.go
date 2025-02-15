package schema

import "errors"

type FloatField struct {
	defaultValue float64
	required     bool
}

func (f *FloatField) Validate(value any) error {
	if _, ok := value.(float64); !ok {
		return errors.New("expected float64")
	}
	return nil
}

func (f *FloatField) Required() Field {
	f.required = true
	return f
}

func (f *FloatField) IsRequired() bool {
	return f.required
}

func (f *FloatField) Default(value any) Field {
	if v, ok := value.(float64); ok {
		f.defaultValue = v
	}
	return f
}

func (f *FloatField) HasDefault() bool {
	return f.defaultValue != 0.0
}

func (f *FloatField) GetDefault() any {
	return f.defaultValue
}

func (f *FloatField) CastTo(target any) (any, error) {
	return target, nil
}

func Float() Field {
	return &FloatField{required: false}
}
