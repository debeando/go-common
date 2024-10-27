package env

import (
	"os"
	"strconv"
)

// Verify exist environment variable.
func Exist(key string) bool {
	if _, ok := os.LookupEnv(key); ok {
		return true
	}
	return false
}

// Retrieve environment variable in String type.
func Get(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Retrieve environment variable in Int type.
func GetInt(key string, fallback int) int {
	if envvalue, ok := os.LookupEnv(key); ok {
		value, _ := strconv.Atoi(envvalue)
		return value
	}

	return fallback
}

// Retrieve environment variable in Int type.
func GetUInt8(key string, fallback uint8) uint8 {
	if envvalue, ok := os.LookupEnv(key); ok {
		value, _ := strconv.Atoi(envvalue)
		return uint8(value)
	}

	return fallback
}

func GetUInt16(key string, fallback uint16) uint16 {
	if envvalue, ok := os.LookupEnv(key); ok {
		value, _ := strconv.Atoi(envvalue)
		return uint16(value)
	}

	return fallback
}

// Retrieve environment variable in Bool type.
func GetBool(key string, fallback bool) bool {
	if envvalue, ok := os.LookupEnv(key); ok {
		value, _ := strconv.ParseBool(envvalue)
		return value
	}

	return fallback
}
