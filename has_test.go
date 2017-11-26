package godash

import "testing"

func TestHas(t *testing.T) {

	HasTests := []struct {
		input  map[string]interface{}
		key    string
		result bool
	}{{
		input:  map[string]interface{}{"test": "test"},
		key:    "test",
		result: true,
	}, {
		input:  map[string]interface{}{"test": "test"},
		key:    "notTest",
		result: false,
	}}

	for _, test := range HasTests {
		result := Has(test.input, test.key)
		if result != test.result {
			t.Fail()
		}
	}
}
