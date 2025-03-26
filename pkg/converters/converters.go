package converters

import (
	"github.com/wk8/go-ordered-map/v2"
	"sort"
)

// ToAnySlice converts a slice of any type to a slice of any type.
func ToAnySlice[T any](in []T) []any {
	out := make([]any, len(in))
	for i, v := range in {
		out[i] = v
	}
	return out
}

// ToOrderedProps converts a map of string to any type to an ordered map of string to any type,
// inserting keys in sorted (alphabetical) order to ensure stable output.
func ToOrderedProps[T any](m map[string]T) *orderedmap.OrderedMap[string, T] {
	om := orderedmap.New[string, T]()

	// Gather all keys
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	// Sort them
	sort.Strings(keys)

	// Insert in sorted order
	for _, k := range keys {
		om.Set(k, m[k])
	}
	return om
}

// ToUint64Ptr converts a uint64 to a uint64 pointer.
func ToUint64Ptr(i uint64) *uint64 {
	return &i
}
