package cqc

import (
	"os"
)

func getEnv(key, fallbackValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallbackValue
}

func getEnvOrPanic(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	panic("environment variable " + key + " is not set")
}
