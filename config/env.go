package config

import (
	"fmt"
	"os"
)

type Config struct {
	PublicHost string
	Port       string

	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

func InitConfig() Config {

	return Config{
		PublicHost: getEnv("DB_HOST", "localhost"),
		Port:       getEnv("DB_PORT", "3306"),
		DBUser:     getEnv("DB_USER", "guser"),
		DBPassword: getEnv("DB_PASSWORD", "pwsd-go"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "localhost"), getEnv("DB_PORT", "3306")),
		DBName:     getEnv("DB_NAME", "golan_api"),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
