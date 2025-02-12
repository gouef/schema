package schema

import (
	"errors"
	"reflect"
)

// ListField for listOf (indexed arrays)
type ListField struct {
	elem Field
}

func (l ListField) Validate(value any) error {
	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Slice {
		return errors.New("expected list (indexed array)")
	}
	// Validate elements in the list
	for i := 0; i < v.Len(); i++ {
		if err := l.elem.Validate(v.Index(i).Interface()); err != nil {
			return errors.New("list element " + string(rune(i)) + ": " + err.Error())
		}
	}
	return nil
}

func ListOf(elem Field) Field {
	return ListField{elem: elem}
}
