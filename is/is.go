package is

// *******************************************************
// **************** Check primitive types ****************
// *******************************************************

// Int checks if the type of the argument is int
func Int(val interface{}) bool {
	switch val.(type) {
	case int:
		return true
	default:
		return false
	}
}

// UInt checks if the type of the argument is uint
func UInt(val interface{}) bool {
	switch val.(type) {
	case uint:
		return true
	default:
		return false
	}
}

// Int8 checks if the type of the argument is int8
func Int8(val interface{}) bool {
	switch val.(type) {
	case int8:
		return true
	default:
		return false
	}
}

// UInt8 checks if the type of the argument is uint8
func UInt8(val interface{}) bool {
	// TODO: should we use reflection to explicitly check if this is int32 or byte (an alias for uint8)
	switch val.(type) {
	case uint8:
		return true
	default:
		return false
	}
}

// Int16 checks if the type of the argument is int16
func Int16(val interface{}) bool {
	switch val.(type) {
	case int16:
		return true
	default:
		return false
	}
}

// UInt16 checks if the type of the argument is uint16
func UInt16(val interface{}) bool {
	switch val.(type) {
	case uint16:
		return true
	default:
		return false
	}
}

// Int32 checks if the type of the argument is int32
func Int32(val interface{}) bool {
	// TODO: should we use reflection to explicitly check if this is int32 or rune (an alias for int32)
	switch val.(type) {
	case int32:
		return true
	default:
		return false
	}
}

// UInt32 checks if the type of the argument is uint32
func UInt32(val interface{}) bool {
	switch val.(type) {
	case uint32:
		return true
	default:
		return false
	}
}

// Int64 checks if the type of the argument is int64
func Int64(val interface{}) bool {
	switch val.(type) {
	case int64:
		return true
	default:
		return false
	}
}

// UInt64 checks if the type of the argument is uint64
func UInt64(val interface{}) bool {
	switch val.(type) {
	case uint64:
		return true
	default:
		return false
	}
}

// Float32 checks if the type of the argument is float32
func Float32(val interface{}) bool {
	switch val.(type) {
	case float32:
		return true
	default:
		return false
	}
}

// Float64 checks if the type of the argument is float64
func Float64(val interface{}) bool {
	switch val.(type) {
	case float64:
		return true
	default:
		return false
	}
}

// Complex64 checks if the type of the argument is complex64
func Complex64(val interface{}) bool {
	switch val.(type) {
	case complex64:
		return true
	default:
		return false
	}
}

// Complex128 checks if the type of the argument is complex128
func Complex128(val interface{}) bool {
	switch val.(type) {
	case complex128:
		return true
	default:
		return false
	}
}

// Bool checks if the type of the argument is bool
func Bool(val interface{}) bool {
	switch val.(type) {
	case bool:
		return true
	default:
		return false
	}
}

// String checks if the type of the argument is string
func String(val interface{}) bool {
	switch val.(type) {
	case string:
		return true
	default:
		return false
	}
}
