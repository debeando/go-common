package env

import (
	"os"
)

// Retrieve environment variable.
func Get(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
