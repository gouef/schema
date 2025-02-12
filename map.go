package schema

import (
	"errors"
	"reflect"
)

// MapField for handling maps
type MapField struct {
	keyType   Field
	valueType Field
}

func (m MapField) Validate(value any) error {
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

func Map(keyType, valueType Field) Field {
	return MapField{keyType: keyType, valueType: valueType}
}
