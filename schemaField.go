package schema

import (
	"strings"
)

type Field interface {
	Validate(value any) error
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
	default:
		return data, nil
	}
}

// Process function with normalization for handling defaults
func Process(schema Field, data any, mergeDefaults bool) error {
	normalizedData, err := Normalize(data)
	if err != nil {
		return err
	}

	// If mergeDefaults is true, you can merge the default values with the provided data
	if mergeDefaults {
		// Add logic for merging defaults if necessary
	}

	return schema.Validate(normalizedData)
}
