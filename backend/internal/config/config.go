package config

import "os"

type Config struct {
	Port           string
	DatabaseURL    string
	Env            string
	AllowedOrigins string
	JWTSecret      string
	PolicyModel    string
	PolicyFile     string
}

func Load() *Config {
	return &Config{
		Port:           getEnv("PORT", "8080"),
		DatabaseURL:    getEnv("DATABASE_URL", ""),
		Env:            getEnv("ENV", "development"),
		AllowedOrigins: getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:3000"),
		JWTSecret:      getEnv("JWT_SECRET", ""),
		PolicyModel:    getEnv("POLICY_MODEL", "./policies/model.conf"),
		PolicyFile:     getEnv("POLICY_FILE", "./policies/policy.csv"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
