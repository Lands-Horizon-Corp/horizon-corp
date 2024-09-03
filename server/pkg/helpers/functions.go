package helpers

import (
	"strconv"
)

func FileSizeStringToInt64(data string) int64 {
	if intValue, err := strconv.ParseInt(data, 10, 64); err == nil {
		return intValue
	}
	return 10 << 20
}
