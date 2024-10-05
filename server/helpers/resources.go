package helpers

func SafeString(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}

func SafeInt(val *int) int {
	if val == nil {
		return 0
	}
	return *val
}
