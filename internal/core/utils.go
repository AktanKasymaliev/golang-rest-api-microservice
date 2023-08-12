package core

import "os"

// Getenv from core give you opportunity for setting default value
func Getenv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}