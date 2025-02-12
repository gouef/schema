package schema

import (
	"errors"
	"reflect"
)

type StructureField struct {
	fields map[string]Field
}

func (s StructureField) Validate(value any) error {
	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Map {
		return errors.New("expected map structure")
	}
	for _, key := range v.MapKeys() {
		field, exists := s.fields[key.String()]
		if !exists {
			return errors.New("unknown field: " + key.String())
		}
		if err := field.Validate(v.MapIndex(key).Interface()); err != nil {
			return errors.New("field " + key.String() + ": " + err.Error())
		}
	}
	return nil
}

func Structure(fields map[string]Field) Field {
	return StructureField{fields: fields}
}

func FromStruct(v any) (Field, error) {
	t := reflect.TypeOf(v)
	if t.Kind() != reflect.Struct {
		return nil, errors.New("expected struct")
	}

	fields := make(map[string]Field)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldName := field.Name
		fieldType := field.Type

		var f Field
		switch fieldType.Kind() {
		case reflect.Bool:
			f = Bool()
		case reflect.Int:
			f = Int()
		case reflect.String:
			f = String()
		case reflect.Float64:
			f = Float()
		default:
			return nil, errors.New("unsupported field type")
		}

		fields[fieldName] = f
	}

	return Structure(fields), nil
}

func CastTo[T any](data map[string]any, v *T) error {
	t := reflect.TypeOf(*v)
	if t.Kind() != reflect.Struct {
		return errors.New("target is not a struct")
	}

	for key, value := range data {
		field := reflect.ValueOf(v).Elem().FieldByName(key)
		if !field.IsValid() {
			return errors.New("unknown field: " + key)
		}
		if field.Kind() == reflect.Int {
			if val, ok := value.(int); ok {
				field.SetInt(int64(val))
			} else {
				return errors.New("type mismatch for field: " + key)
			}
		} else if field.Kind() == reflect.Bool {
			if val, ok := value.(bool); ok {
				field.SetBool(val)
			} else {
				return errors.New("type mismatch for field: " + key)
			}
		} else if field.Kind() == reflect.String {
			if val, ok := value.(string); ok {
				field.SetString(val)
			} else {
				return errors.New("type mismatch for field: " + key)
			}
		} else if field.Kind() == reflect.Float64 {
			if val, ok := value.(float64); ok {
				field.SetFloat(val)
			} else {
				return errors.New("type mismatch for field: " + key)
			}
		} else {
			return errors.New("unsupported field type")
		}
	}

	return nil
}
