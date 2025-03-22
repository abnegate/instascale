package converters

import (
	"github.com/invopop/jsonschema"
	"github.com/wk8/go-ordered-map/v2"
)

func ToAnySlice[T any](in []T) []any {
	out := make([]any, len(in))
	for i, v := range in {
		out[i] = v
	}
	return out
}

func ToOrderedProps(m map[string]*jsonschema.Schema) *orderedmap.OrderedMap[string, *jsonschema.Schema] {
	om := orderedmap.New[string, *jsonschema.Schema]()
	for k, v := range m {
		om.Set(k, v)
	}
	return om
}

func ToUint64Ptr(i uint64) *uint64 {
	return &i
}
