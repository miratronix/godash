package godash

import (
	"errors"
	"strconv"
)

// *******************************************************
// ********** Convert numeric values to string ***********
// *******************************************************

// ToString converts numeric primitives to string
func ToString(num interface{}) (string, error) {
	val := ""

	// Using base 10 for integer-based conversions
	base := 10

	// The formatter 'f' constructs the string without mantissa, i.e. +- d.dddd vs d.ddE+-dd
	floatFormatter := byte('f')

	// The precision controls the number of digits printed by the formatter.
	// -1 uses the smallest number of digits necessary such that ParseFloat will return f exactly.
	floatPrecision := -1

	switch num.(type) {
	case int:
		i := num.(int)
		val = strconv.FormatInt(int64(i), base)
	case int8:
		i8 := num.(int8)
		val = strconv.FormatInt(int64(i8), base)
	case int16:
		i16 := num.(int16)
		val = strconv.FormatInt(int64(i16), base)
	case int32:
		i32 := num.(int32)
		val = strconv.FormatInt(int64(i32), base)
	case int64:
		val = strconv.FormatInt(num.(int64), base)
	case float32:
		f32 := num.(float32)
		val = strconv.FormatFloat(float64(f32), floatFormatter, floatPrecision, 32)
	case float64:
		val = strconv.FormatFloat(num.(float64), floatFormatter, floatPrecision, 64)
	}

	if val == "" {
		return val, errors.New("error converting to string")
	}

	return val, nil
}
