package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func EnvExists(key string) bool {
	_, exists := os.LookupEnv(key)
	return exists
}

func EnvLoad(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}

func EnvInfo() map[string]string {
	environment, _ := godotenv.Unmarshal(strings.Join(os.Environ(), "\n"))
	result := make(map[string]string)
	for key, value := range environment {
		upperKey := strings.ToUpper(key)
		if strings.HasPrefix(upperKey, "APP.") {
			if strings.Contains(upperKey, "PASSWORD") || strings.Contains(upperKey, "SECRET") {
				value = strings.Repeat("*", len(value))
			}
			result[upperKey] = value
		}
	}
	return result
}
