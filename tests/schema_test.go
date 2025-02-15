package tests

import (
	"github.com/gouef/schema"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSchema(t *testing.T) {
	t.Run("ArrayOf", func(t *testing.T) {
		sch := schema.ArrayOf(schema.String())
		normalized, err := schema.Process(sch, []any{"hello", "world"}, false)
		assert.NoError(t, err)
		assert.Equal(t, []any{"hello", "world"}, normalized)

		normalized, err = schema.Process(sch, []any{"hello", 1}, false)
		assert.Error(t, err)
		assert.Equal(t, []any{"hello", 1}, normalized)
	})
	t.Run("ListOf", func(t *testing.T) {
		sch := schema.ListOf(schema.Int())
		normalized, err := schema.Process(sch, []any{1, 2, 3}, false)
		assert.Equal(t, []any{1, 2, 3}, normalized)
		assert.NoError(t, err)
		normalized, err = schema.Process(sch, []any{1, 2, "3"}, false)
		assert.Equal(t, []any{1, 2, "3"}, normalized)
		assert.Error(t, err)
	})
	t.Run("AnyOf", func(t *testing.T) {
		sch := schema.AnyOf(schema.String(), schema.Int())
		normalized, err := schema.Process(sch, "test", false)
		assert.Equal(t, "test", normalized)
		assert.NoError(t, err)
		normalized, err = schema.Process(sch, []any{1, 2, "3"}, false)
		assert.Equal(t, []any{1, 2, "3"}, normalized)
		assert.Error(t, err)
	})

	t.Run("Structure", func(t *testing.T) {
		sch := schema.Structure(map[string]schema.Field{
			"handlers":             schema.AnyOf(schema.ArrayOf(schema.String()), schema.Bool()),
			"processors":           schema.AnyOf(schema.ArrayOf(schema.String()), schema.Bool()),
			"name":                 schema.String().Default("app"),
			"hookToTracy":          schema.Bool().Default(true),
			"tracyBaseUrl":         schema.String(),
			"usePriorityProcessor": schema.Bool().Default(true),
			"accessPriority":       schema.String().Default("INFO"),
			"logDir":               schema.String(),
		})
		data := map[string]any{}
		proc, err := schema.Process(sch, data, true)
		assert.NoError(t, err)
		assert.Equal(t, []any{1, 2, "3"}, proc)
	})
}
