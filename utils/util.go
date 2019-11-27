package utils

import (
	"fmt"
	"os"

	"github.com/segmentio/ksuid"
)

// GetUUID Get UniqueID
func GetUUID() string {
	return fmt.Sprintf("%s", ksuid.New())
}

// GetConfValue Get GetConfValue
func GetConfValue(envKey string) (string, error) {
	if envValue := os.Getenv(envKey); envValue != "" {
		return envValue, nil
	}
	return "", fmt.Errorf("No config found for %s", envKey)
}
