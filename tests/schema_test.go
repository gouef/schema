package tests

import (
	"github.com/gouef/schema"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSchema(t *testing.T) {
	t.Run("ArrayOf", func(t *testing.T) {
		sch := schema.ArrayOf(schema.String())
		err := schema.Process(sch, []any{"hello", "world"}, false)
		assert.NoError(t, err)

		err = schema.Process(sch, []any{"hello", 1}, false)
		assert.Error(t, err)
	})
	t.Run("ListOf", func(t *testing.T) {
		sch := schema.ListOf(schema.Int())
		err := schema.Process(sch, []any{1, 2, 3}, false)
		assert.NoError(t, err)
		err = schema.Process(sch, []any{1, 2, "3"}, false)
		assert.Error(t, err)
	})
	t.Run("AnyOf", func(t *testing.T) {
		sch := schema.AnyOf(schema.String(), schema.Int())
		err := schema.Process(sch, "test", false)
		assert.NoError(t, err)
		err = schema.Process(sch, []any{1, 2, "3"}, false)
		assert.Error(t, err)
	})
}
