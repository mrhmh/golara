package helpers

import "strconv"

func ConvertToInt(value interface{}) int {
	switch value.(type) {
	case int:
		return value.(int)
	case int64:
		return int(value.(int64))
	case float64:
		return int(value.(float64))
	case string:
		i, _ := strconv.Atoi(value.(string))
		return i
	default:
		return 0
	}
}
