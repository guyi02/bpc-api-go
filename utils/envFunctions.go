package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// getEnv retorna o valor de uma key salva no arquivo .env
func GetEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Cannot find this key")
	}
	envValue := os.Getenv(key)
	return envValue
}
