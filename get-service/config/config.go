package config

import (
	"log"
	"os"
)

// Elimina la carga del archivo .env, ya que no es necesario ahora
// func LoadEnv() {
//     if err := godotenv.Load(); err != nil {
//         log.Println("⚠️ No .env file found")
//     }
// }

// Esta función obtiene el valor de una variable de entorno
func GetEnv(key string) string {
    value := os.Getenv(key)
    if value == "" {
        log.Printf("⚠️ Variable de entorno %s no encontrada\n", key)
    }
    return value
}
