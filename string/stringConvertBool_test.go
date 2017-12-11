package string

import (
	"fmt"
	"reflect"
	"testing"
	""
)

func TestErrBool(t *testing.T) {
	ConversionBoolTests := []struct {
		input  string
		error  error
		result bool
	}{{
		input:  "",
		error:  BOOL_ERR,
		result: false,
	}, {
		input:  "True",
		error:  nil,
		result: true,
	}, {
		input:  "true",
		error:  nil,
		result: true,
	}, {
		input:  "False",
		error:  nil,
		result: false,
	}, {
		input:  "false",
		error:  nil,
		result: false,
	}, {
		input:  "1",
		error:  nil,
		result: true,
	}, {
		input:  "0",
		error:  nil,
		result: false,
	}, {
		input:  "TRUE",
		error:  nil,
		result: true,
	}, {
		input:  "FALSE",
		error:  nil,
		result: false,
	}, {
		input:  "TrUe",
		error:  BOOL_ERR,
		result: false,
	}, {
		input:  "FalSe",
		error:  BOOL_ERR,
		result: false,
	}}

	for _, test := range ConversionBoolTests {
		res, err := ToBool(test.input)
		if res != test.result || err != test.error {
			fmt.Println(res, err, reflect.TypeOf(res))
			t.Fail()
		}
	}
}
