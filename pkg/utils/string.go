package utils

func OrString(val, fallback string) string {
	if val != "" {
		return val
	}
	return fallback
}
