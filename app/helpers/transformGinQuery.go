package helpers

import (
	"regexp"
	"strings"
)

// ConvertGinQueryToFilters TODO refactor this function later
func ConvertGinQueryToFilters(queries map[string][]string) map[string]interface{} {

	arrayFields := []string{}

	firstTransform := make(map[string][]string)
	for key, value := range queries {

		if !strings.Contains(key, "[") {
			firstTransform[key] = value
			continue
		}

		// if query is array

		// transform key from "key[] or key[0]" to "key"
		key = regexp.MustCompile("\\[\\d*]").ReplaceAllString(key, "")

		// append key to array fields
		arrayFields = append(arrayFields, key)

		// transform key value from "key[]=[1] and key[]=[2] and key[0]=[3]" to "key=[1,2,3]"
		firstTransform[key] = append(firstTransform[key], value...)
	}

	// transform not array from "key=[val]" to "key=val"
	secondTransform := make(map[string]interface{})
	for key, value := range firstTransform {
		secondTransform[key] = value
		if !InArray(key, arrayFields) {
			secondTransform[key] = value[0]
		}
	}

	return secondTransform
}
