package infra

import (
	"log"
	"os"
	"strconv"
)

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
	}
	return intValue
}
