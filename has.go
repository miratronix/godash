package godash

// Has determines if a key exists in a map
func Has(m map[string]interface{}, key string) bool {
	_, keyExists := m[key]
	return keyExists
}
