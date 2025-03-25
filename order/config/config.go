package config

import (
	"log"
	"os"
	"strconv"
)

// GetEnv reads the environment variable "ENV" and returns its value
func GetEnv() string {
	return getEnvironmentValue("ENV")
}

// GetDataSourceURL reads the environment variable "DATA_SOURCE_URL" and returns its value
func GetDataSourceURL() string {
	return getEnvironmentValue("DATA_SOURCE_URL")
}

// GetApplicationPort reads the environment variable "APPLICATION_PORT" and returns its value as an integer
func GetApplicationPort() int {
	portStr := getEnvironmentValue("APPLICATION_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("APPLICATION_PORT: %s is invalid", portStr)
	}
	return port
}

// getEnvironmentValue reads an environment variable by key and logs a fatal error if missing
func getEnvironmentValue(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s environment variable is missing.", key)
	}
	return value
}
