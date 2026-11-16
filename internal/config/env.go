package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv(env string) {
	err := godotenv.Load(env)
	if err != nil {
		log.Fatalf("error loadEnv: no environment file (.env) was found. The systems environment variables are used: %v", err)
	}
}
