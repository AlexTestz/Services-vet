package config

import (
	"log"
	"os"
)

// Remove the .env file load, as it is no longer necessary
// func LoadEnv() {
//     if err := godotenv.Load(); err != nil {
//         log.Println(“⚠️ No .env file found”)
//     }
// }

// This function obtains the value of an environment variable
func GetEnv(key string) string {
    value := os.Getenv(key)
    if value == "" {
        log.Printf("⚠️ Variable de entorno %s no encontrada\n", key)
    }
    return value
}
