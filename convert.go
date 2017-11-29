package godash

import (
	"strconv"
)

type AnyType interface {
	convert()
}

// *******************************************************
// ************** Convert string to integer **************
// *******************************************************

func ToInt(str string) (int, error) {
	// Covert a string to a base-10 integer
	val, err := strconv.ParseInt(str, 0, 0)
	if err != nil {
		// This error handling also takes care of overflows. If string is out of range for the int type,
		// strconv will throw an "ErrRange" error, indicating the overflow.
		return 0, INT_ERR
	}
	return int(val), nil
}

func ToInt8(str string) (int8, error) {
	// Base "0" forces strconv.ParseInt() to infer the base
	val, err := strconv.ParseInt(str, 0, 8)
	if err != nil {
		return 0, INT_ERR
	}

	// strconv.ParseInt() returns int64, so cast to int8
	// If overflow error occurs, it will be handled above
	return int8(val), nil
}

func ToInt16(str string) (int16, error) {
	val, err := strconv.ParseInt(str, 0, 16)
	if err != nil {
		return 0, INT_ERR
	}
	return int16(val), nil
}

func ToInt32(str string) (int32, error) {
	val, err := strconv.ParseInt(str, 0, 32)
	if err != nil {
		return 0, INT_ERR
	}
	return int32(val), nil
}

func ToInt64(str string) (int64, error) {
	val, err := strconv.ParseInt(str, 0, 64)
	if err != nil {
		return 0, INT_ERR
	}
	return val, nil
}

// *******************************************************
// ************** Convert string to floats ***************
// *******************************************************

func ToFloat32(str string) (float32, error) {
	val, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0, FLOAT_ERR
	}
	return float32(val), nil
}

func ToFloat64(str string) (float64, error) {
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, FLOAT_ERR
	}
	return val, nil
}

// *******************************************************
// ************** Convert string to boolean **************
// *******************************************************

// Converts "1,  True,  true" to true
//          "0, False, false" to false
func ToBool(str string) (bool, error) {
	val, err := strconv.ParseBool(str)
	if err != nil {
		return false, BOOL_ERR
	}
	return val, nil
}

// *******************************************************
// ********** Convert numeric values to string ***********
// *******************************************************

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
		return val, STR_ERR
	}

	return val, nil
}
