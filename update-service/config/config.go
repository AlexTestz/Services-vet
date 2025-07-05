package config

import (
	"os"
	//	"github.com/joho/godotenv"
)

//func LoadEnv() {
//	if err := godotenv.Load(); err != nil {
//		log.Println("No .env file found")
//	}
//}

func GetEnv(key string) string {
	return os.Getenv(key)
}
// GetEnvOrDefault returns the value of the environment variable or a default value if not set