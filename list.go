package schema

import (
	"errors"
	"reflect"
)

// ListField for listOf (indexed arrays)
type ListField struct {
	elem         Field
	defaultValue []any
}

func (l *ListField) Validate(value any) error {
	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Slice {
		return errors.New("expected list (indexed array)")
	}
	for i := 0; i < v.Len(); i++ {
		if err := l.elem.Validate(v.Index(i).Interface()); err != nil {
			return errors.New("list element " + string(rune(i)) + ": " + err.Error())
		}
	}
	return nil
}

func (l *ListField) Default(value any) Field {
	if v, ok := value.([]any); ok {
		l.defaultValue = v
	}
	return l
}

func (l *ListField) HasDefault() bool {
	return len(l.defaultValue) > 0
}

func (l *ListField) GetDefault() any {
	return l.defaultValue
}

func (l *ListField) CastTo(target any) (any, error) {
	return target, nil
}

func ListOf(elem Field) Field {
	return &ListField{elem: elem}
}
