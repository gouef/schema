package schema

import "errors"

type StringField struct {
	defaultValue string
	required     bool
}

func (s *StringField) Validate(value any) error {
	if _, ok := value.(string); !ok {
		return errors.New("expected string")
	}
	return nil
}

func (s *StringField) Required() Field {
	s.required = true
	return s
}

func (s *StringField) IsRequired() bool {
	return s.required
}

func (s *StringField) Default(value any) Field {
	if v, ok := value.(string); ok {
		s.defaultValue = v
	}
	return s
}

func (s *StringField) HasDefault() bool {
	return len(s.defaultValue) > 0
}

func (s *StringField) GetDefault() any {
	return s.defaultValue
}

func (s *StringField) CastTo(target any) (any, error) {
	return target, nil
}

func String() Field {
	return &StringField{required: false}
}
