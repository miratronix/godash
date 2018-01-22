package string

import (
	"strconv"
)

// *******************************************************
// ************** Convert string to boolean **************
// *******************************************************

// ToBool converts a string to bool
// Converts "1,  True,  true" to true
//          "0, False, false" to false
func ToBool(str string) (bool, error) {
	val, err := strconv.ParseBool(str)
	if err != nil {
		return false, BOOL_ERR
	}
	return val, nil
}