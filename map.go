package schema

import (
	"errors"
	"reflect"
)

// MapField for handling maps
type MapField struct {
	keyType      Field
	valueType    Field
	defaultValue map[any]any
	required     bool
}

func (m *MapField) Validate(value any) error {
	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Map {
		return errors.New("expected map")
	}
	for _, key := range v.MapKeys() {
		if err := m.keyType.Validate(key.Interface()); err != nil {
			return errors.New("map key: " + err.Error())
		}
		if err := m.valueType.Validate(v.MapIndex(key).Interface()); err != nil {
			return errors.New("map value: " + err.Error())
		}
	}
	return nil
}

func (m *MapField) Required() Field {
	m.required = true
	return m
}

func (m *MapField) IsRequired() bool {
	return m.required
}

func (m *MapField) Default(value any) Field {
	if v, ok := value.(map[any]any); ok {
		m.defaultValue = v
	}
	return m
}

func (m *MapField) HasDefault() bool {
	return len(m.defaultValue) > 0
}

func (m *MapField) GetDefault() any {
	return m.defaultValue
}

func (m *MapField) CastTo(target any) (any, error) {
	return target, nil
}

func Map(keyType, valueType Field) Field {
	return &MapField{keyType: keyType, valueType: valueType, required: false}
}
