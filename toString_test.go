package godash

import (
	"fmt"
	"reflect"
	"testing"
)

func TestErrString(t *testing.T) {
	var i int = 1000
	var i8 int8 = 8
	var i16 int16 = 16
	var i32 int32 = 32
	var i64 int64 = 64
	var f32 float32 = 3.2
	var f64 float64 = 6.4

	ConversionStringTests := []struct {
		input  interface{}
		error  error
		result string
	}{{
		input:  i,
		error:  nil,
		result: "1000",
	}, {
		input:  -1 * i,
		error:  nil,
		result: "-1000",
	}, {
		input:  i8,
		error:  nil,
		result: "8",
	}, {
		input:  -1 * i8,
		error:  nil,
		result: "-8",
	}, {
		input:  i16,
		error:  nil,
		result: "16",
	}, {
		input:  -1 * i16,
		error:  nil,
		result: "-16",
	}, {
		input:  i32,
		error:  nil,
		result: "32",
	}, {
		input:  -1 * i32,
		error:  nil,
		result: "-32",
	}, {
		input:  i64,
		error:  nil,
		result: "64",
	}, {
		input:  -1 * i64,
		error:  nil,
		result: "-64",
	}, {
		input:  i64,
		error:  nil,
		result: "64",
	}, {
		input:  f32,
		error:  nil,
		result: "3.2",
	}, {
		input:  -1 * f32,
		error:  nil,
		result: "-3.2",
	}, {
		input:  f64,
		error:  nil,
		result: "6.4",
	}, {
		input:  -1 * f64,
		error:  nil,
		result: "-6.4",
	}}

	for _, test := range ConversionStringTests {
		res, err := ToString(test.input)
		if res != test.result || err != test.error {
			fmt.Println(res, err, reflect.TypeOf(res))
			t.Fail()
		}
	}
}
