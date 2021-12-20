package envvar

import (
	"log"
	"os"
	"strconv"
)

// ProtocolDecimalPlacesAfterDateEnv Const to register the environment name
const ProtocolDecimalPlacesAfterDateEnv = "PROTOCOL_DECIMAL_PLACES_AFTER_DATE"

// RedisAddressEnv Const to register the environment name
const RedisAddressEnv = "REDIS_ADDRESS"

// GetIntegerEnv Get the integer value or default from environment
func GetIntegerEnv(envName string, defaultValue int) int {
	envValue := os.Getenv(envName)
	if envValue == "" {
		log.Printf("Env %s not defined. Using default value: %d", envName, defaultValue)
		return defaultValue
	}
	intValue, converterError := strconv.Atoi(envValue)
	if converterError != nil {
		log.Printf("Error to convert env %s with value: %s to integer value. Using default value: %d",
			envName, envValue, defaultValue)
		return defaultValue
	}
	return intValue
}

// GetStringEnv Get the string value or default from environment
func GetStringEnv(envName string, defaultValue string) string {
	envValue := os.Getenv(envName)
	if envValue == "" {
		log.Printf("Env %s not defined. Using default value: %s", envName, defaultValue)
		return defaultValue
	}
	return envValue
}
