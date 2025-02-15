package schema

import (
	"errors"
	"fmt"
	"strings"
)

type Field interface {
	Validate(value any) error
	Default(value any) Field
	CastTo(target any) (any, error)
	HasDefault() bool
	GetDefault() any
	Required() Field
	IsRequired() bool
}

// Normalize function to clean and prepare data
func Normalize(data any) (any, error) {
	switch v := data.(type) {
	case string:
		return strings.TrimSpace(v), nil
	case []string:
		// Normalize string slices
		for i, val := range v {
			v[i] = strings.TrimSpace(val)
		}
		return v, nil
	case []int:
		// Example normalization for int slices
		return v, nil
	case map[string]any:
		for k, val := range v {
			normalizedVal, err := Normalize(val)
			if err != nil {
				return nil, err
			}
			v[k] = normalizedVal
		}
		return v, nil
	default:
		return data, nil
	}
}

// Process function with normalization for handling defaults
func Process(schema Field, data any, mergeDefaults bool) (any, error) {
	normalizedData, err := Normalize(data)
	if err != nil {
		return nil, err
	}

	// If mergeDefaults is true, you can merge the default values with the provided data
	if mergeDefaults {
		// Add logic for merging defaults if necessary
		if structure, ok := schema.(*StructureField); ok {
			for key, field := range structure.fields {
				if _, exists := normalizedData.(map[string]any)[key]; !exists && field.HasDefault() {
					normalizedData.(map[string]any)[key] = field.GetDefault()
				} else if !exists && field.IsRequired() {
					return nil, errors.New(fmt.Sprintf("%s is required.", key))
				}
			}
		}
	}

	return normalizedData, schema.Validate(normalizedData)
}
