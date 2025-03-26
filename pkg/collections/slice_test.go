package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		item     int
		expected bool
	}{
		{
			name:     "Empty slice",
			slice:    []int{},
			item:     5,
			expected: false,
		},
		{
			name:     "Item not in slice",
			slice:    []int{1, 2, 3},
			item:     4,
			expected: false,
		},
		{
			name:     "Item in slice",
			slice:    []int{1, 2, 3},
			item:     3,
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := contains(tc.slice, tc.item)
			assert.Equal(t, tc.expected, result)
		})
	}
}
