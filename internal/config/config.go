package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}
	return nil
}

func CountWithPrefix(prefix string) int {
	envs := GetWithPrefix(prefix)
	return len(envs)
}

func GetWithPrefix(prefix string, secrets ...string) map[string]string {
	environment, _ := godotenv.Unmarshal(strings.Join(os.Environ(), "\n"))
	result := make(map[string]string)
	for key, value := range environment {
		upperKey := strings.ToUpper(key)
		if strings.HasPrefix(upperKey, strings.ToUpper(prefix)) {
			for _, s := range secrets {
				upperSecret := strings.ToUpper(s)
				if strings.Contains(upperKey, upperSecret) || strings.Contains(upperSecret, upperKey) {
					value = strings.Repeat("*", len(value))
				}
			}
			result[upperKey] = value
		}
	}
	return result
}
