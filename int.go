package schema

import "errors"

type IntField struct {
	defaultValue int
}

func (i *IntField) Validate(value any) error {
	if _, ok := value.(int); !ok {
		return errors.New("expected integer")
	}
	return nil
}

func (i *IntField) Default(value any) Field {
	if v, ok := value.(int); ok {
		i.defaultValue = v
	}
	return i
}

func (i *IntField) HasDefault() bool {
	return i.defaultValue != 0
}

func (i *IntField) GetDefault() any {
	return i.defaultValue
}

func (i *IntField) CastTo(target any) (any, error) {
	return target, nil
}

func Int() Field {
	return &IntField{}
}
