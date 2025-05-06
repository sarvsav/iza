package dbstore

import (
	"fmt"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// toString safely converts an interface{} to string
func toString(val interface{}) string {
	switch v := val.(type) {
	case string:
		return v
	case primitive.ObjectID:
		return v.Hex()
	case fmt.Stringer:
		return v.String()
	case []byte:
		return string(v)
	case nil:
		return ""
	default:
		return fmt.Sprintf("%v", v)
	}
}

// toInt64 safely converts an interface{} to int64
func toInt64(val interface{}) int64 {
	switch v := val.(type) {
	case int64:
		return v
	case int32:
		return int64(v)
	case int:
		return int64(v)
	case float64:
		return int64(v)
	case float32:
		return int64(v)
	case string:
		if i, err := strconv.ParseInt(v, 10, 64); err == nil {
			return i
		}
	case primitive.DateTime:
		return v.Time().Unix()
	}
	return 0
}

// toBool safely converts an interface{} to bool
func toBool(val interface{}) bool {
	switch v := val.(type) {
	case bool:
		return v
	case string:
		if v == "true" || v == "1" {
			return true
		}
	case int, int64, int32:
		return toInt64(v) != 0
	}
	return false
}
