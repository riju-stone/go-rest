package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	Host       string
	Port       string
	DBUser     string
	DBPassword string
	DBAddr     string
	DBName     string
}

var EnvData = initConfig()

func initConfig() EnvConfig {
	error := godotenv.Load(".env.local")

	if error != nil {
		log.Fatal(error)
	}

	return EnvConfig{
		Host:       getEnvVars("PUBLIC_HOST", "http://localhost"),
		Port:       getEnvVars("PORT", "8080"),
		DBUser:     getEnvVars("DB_USER", "root"),
		DBPassword: getEnvVars("DB_PASSWORD", "password"),
		DBAddr:     getEnvVars("DB_URL", "127.0.0.1"),
		DBName:     getEnvVars("DB_NAME", "sample"),
	}
}

func getEnvVars(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
