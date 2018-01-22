package godash

// Defaults recursively merges maps from left to right, without overriding existing keys
func Defaults(maps ...map[string]interface{}) map[string]interface{} {

	// Initialize a return map
	returnMap := map[string]interface{}{}

	for _, m := range maps {
		for key, value := range m {

			// If we don't already have the key in the return map, add it
			if !Has(returnMap, key) {
				returnMap[key] = value
				continue
			}

			returnMapValue, castReturnMap := returnMap[key].(map[string]interface{})
			mValue, castM := m[key].(map[string]interface{})

			// Not a map in both, ignore it
			if !castReturnMap || !castM {
				continue
			}

			// Both are maps, keep merging
			returnMap[key] = Defaults(returnMapValue, mValue)
		}
	}

	return returnMap
}
