package schema

import (
	"github.com/ksensehq/eventnative/test"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMap(t *testing.T) {
	tests := []struct {
		name           string
		mappings       []string
		inputObject    map[string]interface{}
		expectedObject map[string]interface{}
	}{
		{
			"nil input object",
			nil,
			nil,
			nil,
		},
		{
			"Empty mappings and input object",
			nil,
			map[string]interface{}{},
			map[string]interface{}{},
		},
		{
			"Dummy mapper doesn't change input json",
			nil,
			map[string]interface{}{
				"key1": map[string]interface{}{
					"subkey1": 123,
				},
				"key2": "value",
			},
			map[string]interface{}{
				"key1": map[string]interface{}{
					"subkey1": 123,
				},
				"key2": "value",
			},
		},
		{
			"Map unflatten object",
			[]string{"/key1 -> /key10", "/key2/subkey2-> /key11", "/key4/subkey1 ->", "/key4/subkey3 ->",
				"/key4/subkey4 -> /key4", "/key5 -> /key6/subkey1", "/key3/subkey1 -> /key7", "/key3 -> /key2/subkey1"},
			map[string]interface{}{
				"key1": map[string]interface{}{
					"subkey1": map[string]interface{}{
						"subsubkey1": 123,
						"subsubkey2": 123,
					},
				},
				"key2": "value",
				"key3": 999,
				"key4": map[string]interface{}{
					"subkey1": map[string]interface{}{
						"subsubkey1": 123,
						"subsubkey2": 123,
					},
					"subkey2": 123,
				},
				"key5": 888,
			},
			map[string]interface{}{
				"key10": map[string]interface{}{
					"subkey1": map[string]interface{}{
						"subsubkey1": 123,
						"subsubkey2": 123,
					},
				},
				"key2": "value",
				"key3": 999,
				"key4": map[string]interface{}{
					"subkey2": 123,
				},
				"key6": map[string]interface{}{
					"subkey1": 888,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapper, _, err := NewFieldMapper(Default, tt.mappings)
			require.NoError(t, err)

			actualObject, _ := mapper.Map(tt.inputObject)
			test.ObjectsEqual(t, tt.expectedObject, actualObject, "Mapped objects aren't equal")
		})
	}
}

func TestStrictMap(t *testing.T) {
	tests := []struct {
		name           string
		mappings       []string
		inputObject    map[string]interface{}
		expectedObject map[string]interface{}
	}{
		{
			"nil input object",
			nil,
			nil,
			nil,
		},
		{
			"Empty mappings and input object",
			nil,
			map[string]interface{}{},
			map[string]interface{}{},
		},
		{
			"Dummy mapper doesn't change input json",
			nil,
			map[string]interface{}{
				"key1": map[string]interface{}{
					"subkey1": 123,
				},
				"key2": "value",
			},
			map[string]interface{}{
				"key1": map[string]interface{}{
					"subkey1": 123,
				},
				"key2": "value",
			},
		},
		{
			"Map unflatten object",
			[]string{"/key1 -> /key10", "/key2/subkey2-> /key11", "/key4/subkey1 ->", "/key4/subkey3 ->",
				"/key4/subkey4 -> /key4", "/key5 -> /key6/subkey1", "/key3/subkey1 -> /key7", "/key3 -> /key2/subkey1"},
			map[string]interface{}{
				"key1": map[string]interface{}{
					"subkey1": map[string]interface{}{
						"subsubkey1": 123,
						"subsubkey2": 123,
					},
				},
				"key2": "value",
				"key3": 999,
				"key4": map[string]interface{}{
					"subkey1": map[string]interface{}{
						"subsubkey1": 123,
						"subsubkey2": 123,
					},
					"subkey2": 123,
				},
				"key5": 888,
			},
			map[string]interface{}{
				"key10": map[string]interface{}{
					"subkey1": map[string]interface{}{
						"subsubkey1": 123,
						"subsubkey2": 123,
					},
				},
				"key2": map[string]interface{}{
					"subkey1": 999,
				},
				"key6": map[string]interface{}{
					"subkey1": 888,
				},
			},
		},
		{
			"Minify object test",
			[]string{"/key1 -> /key10", "/key2-> /key11", "/key3->/key12"},
			map[string]interface{}{
				"key1": map[string]interface{}{
					"subkey1": map[string]interface{}{
						"subsubkey1": 123,
						"subsubkey2": 123,
					},
				},
				"key2": "value",
				"key3": 999,
				"key4": map[string]interface{}{
					"subkey1": map[string]interface{}{
						"subsubkey1": 123,
						"subsubkey2": 123,
					},
					"subkey2": 123,
				},
				"key5": 888,
			},
			map[string]interface{}{
				"key10": map[string]interface{}{
					"subkey1": map[string]interface{}{
						"subsubkey1": 123,
						"subsubkey2": 123,
					},
				},
				"key11": "value",
				"key12": 999,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapper, _, err := NewFieldMapper(Strict, tt.mappings)
			require.NoError(t, err)

			actualObject, _ := mapper.Map(tt.inputObject)
			test.ObjectsEqual(t, tt.expectedObject, actualObject, "Mapped objects aren't equal")
		})
	}
}
