package godash

import (
	"fmt"
	"reflect"
	"testing"
)

func TestErrAllConversions(test *testing.T) {
	TestErrInt(test)
	TestErrInt8(test)
	TestErrInt16(test)
	TestErrInt32(test)
	TestErrInt64(test)
	TestErrBool(test)
	TestErrString(test)
}

func TestErrInt(t *testing.T) {
	ConversionIntTests := []struct {
		input  string
		error  error
		result int
	}{{
		input:  "",
		error:  INT_ERR,
		result: 0,
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
		input:  "9223372036854775808",
		error:  INT_ERR,
		result: 0,
	}}

	for _, test := range ConversionIntTests {
		res, err := ToInt(test.input)
		if res != test.result || err != test.error {
			fmt.Println(res, err, reflect.TypeOf(res))
			t.Fail()
		}
	}
}

func TestErrInt8(t *testing.T) {
	ConversionInt8Tests := []struct {
		input  string
		error  error
		result int8
	}{{
		input:  "",
		error:  INT_ERR,
		result: 0,
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
		input:  "128",
		error:  INT_ERR,
		result: 0,
	}}

	for _, test := range ConversionInt8Tests {
		res, err := ToInt8(test.input)
		if res != test.result || err != test.error {
			fmt.Println(res, err, reflect.TypeOf(res))
			t.Fail()
		}
	}
}

func TestErrInt16(t *testing.T) {
	ConversionInt16Tests := []struct {
		input  string
		error  error
		result int16
	}{{
		input:  "",
		error:  INT_ERR,
		result: 0,
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
		input:  "32769",
		error:  INT_ERR,
		result: 0,
	}}

	for _, test := range ConversionInt16Tests {
		res, err := ToInt16(test.input)
		if res != test.result || err != test.error {
			fmt.Println(res, err, reflect.TypeOf(res))
			t.Fail()
		}
	}
}

func TestErrInt32(t *testing.T) {
	ConversionInt32Tests := []struct {
		input  string
		error  error
		result int32
	}{{
		input:  "",
		error:  INT_ERR,
		result: 0,
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
		input:  "2147483648",
		error:  INT_ERR,
		result: 0,
	}}

	for _, test := range ConversionInt32Tests {
		res, err := ToInt32(test.input)
		if res != test.result || err != test.error {
			fmt.Println(res, err, reflect.TypeOf(res))
			t.Fail()
		}
	}
}

func TestErrInt64(t *testing.T) {
	ConversionInt64Tests := []struct {
		input  string
		error  error
		result int64
	}{{
		input:  "",
		error:  INT_ERR,
		result: 0,
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
		input:  "9223372036854775808",
		error:  INT_ERR,
		result: 0,
	}}

	for _, test := range ConversionInt64Tests {
		res, err := ToInt64(test.input)
		if res != test.result || err != test.error {
			fmt.Println(res, err, reflect.TypeOf(res))
			t.Fail()
		}
	}
}

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

func TestErrString(t *testing.T) {
	var i int16 = 1000
	var i8 int = 8
	var i16 int = 16
	var i32 int = 32
	var i64 int = 64
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
