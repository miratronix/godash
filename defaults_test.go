package godash

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDefaults(t *testing.T) {

	DefaultsTests := []struct {
		args   []map[string]interface{}
		result map[string]interface{}
	}{{
		args: []map[string]interface{}{
			map[string]interface{}{"test1": "test1"},
			map[string]interface{}{"test2": "test2"},
		},
		result: map[string]interface{}{
			"test1": "test1",
			"test2": "test2",
		},
	}, {
		args: []map[string]interface{}{
			map[string]interface{}{"test1": "test1"},
			map[string]interface{}{"test1": "test2"},
		},
		result: map[string]interface{}{
			"test1": "test1",
		},
	}, {
		args: []map[string]interface{}{
			map[string]interface{}{"map": map[string]interface{}{"test1": "test1"}},
			map[string]interface{}{"map": map[string]interface{}{"test2": "test2"}},
		},
		result: map[string]interface{}{
			"map": map[string]interface{}{"test1": "test1", "test2": "test2"},
		},
	}, {
		args: []map[string]interface{}{
			map[string]interface{}{"map": map[string]interface{}{"test1": "test1"}},
			map[string]interface{}{"map": map[string]interface{}{"test1": "test2"}},
		},
		result: map[string]interface{}{
			"map": map[string]interface{}{"test1": "test1"},
		},
	}, {
		args: []map[string]interface{}{
			map[string]interface{}{"map": map[string]interface{}{"test1": "test1"}},
			map[string]interface{}{"map": "notMap"},
		},
		result: map[string]interface{}{
			"map": map[string]interface{}{"test1": "test1"},
		},
	}, {
		args: []map[string]interface{}{
			map[string]interface{}{"map": map[string]interface{}{"test1": "test1"}},
			map[string]interface{}{"map": "notMap"},
			map[string]interface{}{"map": map[string]interface{}{"test3": "test3"}},
		},
		result: map[string]interface{}{
			"map": map[string]interface{}{"test1": "test1", "test3": "test3"},
		},
	}}

	for i, test := range DefaultsTests {
		result := Defaults(test.args...)

		if !reflect.DeepEqual(result, test.result) {
			expected, _ := json.Marshal(test.result)
			actual, _ := json.Marshal(result)
			t.Errorf("Failed on test Defaults.%d:\n Expected: %s\n Actual: %s\n", i, string(expected), string(actual))
		}
	}
}
