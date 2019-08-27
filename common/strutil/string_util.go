package strutil

func IsNull(value, defaultValue string) string {
	if len(value) > 0 {
		return value
	}

	return defaultValue
}
