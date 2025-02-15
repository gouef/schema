package schema

import (
	"errors"
	"reflect"
)

type StructureField struct {
	fields       map[string]Field
	defaultValue map[string]any
	required     bool
}

func (s *StructureField) Validate(value any) error {
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

func (s *StructureField) Required() Field {
	s.required = true
	return s
}

func (s *StructureField) IsRequired() bool {
	return s.required
}

func (s *StructureField) Default(value any) Field {
	if v, ok := value.(map[string]any); ok {
		s.defaultValue = v
	}
	return s
}

func (s *StructureField) CastTo(target any) (any, error) {
	return target, nil
}

func Structure(fields map[string]Field) Field {
	return &StructureField{fields: fields, required: false}
}

func (s *StructureField) HasDefault() bool {
	return len(s.defaultValue) > 0
}

func (s *StructureField) GetDefault() any {
	return s.defaultValue
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
		case reflect.Slice:
			elemType := fieldType.Elem()
			if elemType.Kind() == reflect.String {
				f = Array(String())
			} else if elemType.Kind() == reflect.Int {
				f = Array(Int())
			} else {
				return nil, errors.New("unsupported array element type")
			}
		case reflect.Map:
			keyType := fieldType.Key()
			valueType := fieldType.Elem()
			if keyType.Kind() == reflect.String && valueType.Kind() == reflect.Int {
				f = Map(String(), Int())
			} else {
				return nil, errors.New("unsupported map key or value type")
			}
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
