package string

import (
	"strconv"
)

// *******************************************************
// ************** Convert string to integer **************
// *******************************************************

// ToInt converts a string to int
func ToInt(str string) (int, error) {
	// Covert a string to a base-10 integer
	// The second argument is the base of the string to be converted. If it 0, then ParseInt()
	// infers the base.
	val, err := strconv.ParseInt(str, 0, 0)
	if err != nil {
		// This error handling also takes care of overflows. If string is out of range for the int type,
		// strconv will throw an "ErrRange" error, indicating the overflow.
		return 0, INT_ERR
	}
	return int(val), nil
}

// ToInt8 converts a string to int8
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

// ToInt16 converts a string to int16
func ToInt16(str string) (int16, error) {
	val, err := strconv.ParseInt(str, 0, 16)
	if err != nil {
		return 0, INT_ERR
	}
	return int16(val), nil
}

// ToInt32 converts a string to int32
func ToInt32(str string) (int32, error) {
	val, err := strconv.ParseInt(str, 0, 32)
	if err != nil {
		return 0, INT_ERR
	}
	return int32(val), nil
}

// ToInt64 converts a string to int32
func ToInt64(str string) (int64, error) {
	val, err := strconv.ParseInt(str, 0, 64)
	if err != nil {
		return 0, INT_ERR
	}
	return val, nil
}
