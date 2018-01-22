package string

import (
	"fmt"
	"reflect"
	"testing"
)

func TestErrAllIntConversions(test *testing.T) {
	TestErrInt(test)
	TestErrInt8(test)
	TestErrInt16(test)
	TestErrInt32(test)
	TestErrInt64(test)
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