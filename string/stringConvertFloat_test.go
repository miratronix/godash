package string

import (
	"fmt"
	"reflect"
	"testing"
)

func TestErrAllFloatConversions(test *testing.T) {
	TestErrFloat32(test)
	TestErrFloat64(test)
}

func TestErrFloat32(t *testing.T) {
	ConversionInt32Tests := []struct {
		input  string
		error  error
		result float32
	}{{
		input:  "",
		error:  FLOAT_ERR,
		result: 0,
	}, {
		input:  "11.11",
		error:  nil,
		result: 11.11,
	}, {
		input:  "+1.1",
		error:  nil,
		result: 1.1,
	}, {
		input:  "-2.71828",
		error:  nil,
		result: -2.71828,
	}, {
		input:  "100",
		error:  nil,
		result: 100,
	}, {
		input:  "+10",
		error:  nil,
		result: 10,
	}, {
		input:  "-5",
		error:  nil,
		result: -5,
	}, {
		input:  "3.142e2",
		error:  nil,
		result: 314.2,
	}, {
		input:  "-3.142e2",
		error:  nil,
		result: -314.2,
	}, {
		input:  "3.142e2",
		error:  nil,
		result: 3.142e2,
	}}

	for _, test := range ConversionInt32Tests {
		res, err := ToFloat32(test.input)
		if res != test.result || err != test.error {
			fmt.Println(res, err, reflect.TypeOf(res))
			t.Fail()
		}
	}
}

func TestErrFloat64(t *testing.T) {
	ConversionInt64Tests := []struct {
		input  string
		error  error
		result float64
	}{{
		input:  "",
		error:  FLOAT_ERR,
		result: 0,
	}, {
		input:  "11.11",
		error:  nil,
		result: 11.11,
	}, {
		input:  "+1.1",
		error:  nil,
		result: 1.1,
	}, {
		input:  "-2.71828",
		error:  nil,
		result: -2.71828,
	}, {
		input:  "100",
		error:  nil,
		result: 100,
	}, {
		input:  "+10",
		error:  nil,
		result: 10,
	}, {
		input:  "-5",
		error:  nil,
		result: -5,
	}, {
		input:  "3.142e2",
		error:  nil,
		result: 314.2,
	}, {
		input:  "-3.142e2",
		error:  nil,
		result: -314.2,
	}, {
		input:  "3.142e2",
		error:  nil,
		result: 3.142e2,
	}}

	for _, test := range ConversionInt64Tests {
		res, err := ToFloat64(test.input)
		if res != test.result || err != test.error {
			fmt.Println(test.result, test.error, reflect.TypeOf(res))
			fmt.Println("ERROR: ", res, err, reflect.TypeOf(res))
			t.Fail()
		}
	}
}
