package string

import (
	"strconv"
)

// *******************************************************
// ************** Convert string to floats ***************
// *******************************************************

// ToFloat32 converts a string to float32
func ToFloat32(str string) (float32, error) {
	val, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0, FLOAT_ERR
	}
	return float32(val), nil
}

// ToFloat64 converts a string to float64
func ToFloat64(str string) (float64, error) {
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, FLOAT_ERR
	}
	return val, nil
}
