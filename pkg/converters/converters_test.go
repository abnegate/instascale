package converters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToAnySlice(t *testing.T) {
	t.Run("Empty slice", func(t *testing.T) {
		var in []int
		out := ToAnySlice(in)
		assert.NotNil(t, out)
		assert.Len(t, out, 0, "Expected empty slice")
	})

	t.Run("Non-empty slice", func(t *testing.T) {
		in := []string{"alpha", "beta"}
		out := ToAnySlice(in)
		assert.Len(t, out, 2)
		assert.Equal(t, "alpha", out[0])
		assert.Equal(t, "beta", out[1])
	})
}

func TestToOrderedProps(t *testing.T) {
	t.Run("Empty map", func(t *testing.T) {
		in := map[string]string{}
		out := ToOrderedProps(in)
		assert.NotNil(t, out)
		assert.Equal(t, 0, out.Len(), "Expected empty ordered map")
	})

	t.Run("Non-empty map", func(t *testing.T) {
		in := map[string]int{
			"one": 1,
			"two": 2,
		}
		out := ToOrderedProps(in)
		assert.NotNil(t, out)
		assert.Equal(t, 2, out.Len(), "Expected ordered map with 2 entries")

		val1, ok1 := out.Get("one")
		assert.True(t, ok1)
		assert.Equal(t, 1, val1)

		val2, ok2 := out.Get("two")
		assert.True(t, ok2)
		assert.Equal(t, 2, val2)
	})
}

func TestToUint64Ptr(t *testing.T) {
	t.Run("Converts uint64 to pointer", func(t *testing.T) {
		var val uint64 = 42
		ptr := ToUint64Ptr(val)
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}
