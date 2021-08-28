package utils

// AssertParamType returns true if the value is string and false if map, and return the value in it's type
func AssertTypeIsString(value interface{}) (bool, *string, *map[string]string) {
	if s, ok := value.(string); ok {
		return true, &s, nil
	}

	if m, ok := value.(map[interface{}]interface{}); ok {
		mapOfStrings := make(map[string]string)
		for k, v := range m {
			mapOfStrings[k.(string)] = v.(string)
		}
		return false, nil, &mapOfStrings
	}
	panic("can not assert instractions values")

	// return nil, nil, nil, errors.New("")
}
