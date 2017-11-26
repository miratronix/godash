package godash

// Determines if a key exists in the map
func Has(m map[string]interface{}, key string) bool {
	_, keyExists := m[key]
	return keyExists
}
