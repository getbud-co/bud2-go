package config

import "os"

type Config struct {
	Port            string
	DatabaseURL     string
	Env             string
	AllowedOrigins  string
	JWTSecret       string
	PolicyModel     string
	PolicyFile      string
	OTelEndpoint    string
	OTelServiceName string
	OTelEnvironment string
}

func Load() *Config {
	cfg := &Config{
		Port:            getEnv("PORT", "8080"),
		DatabaseURL:     getEnv("DATABASE_URL", ""),
		Env:             getEnv("ENV", "development"),
		AllowedOrigins:  getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:3000"),
		JWTSecret:       getEnv("JWT_SECRET", ""),
		PolicyModel:     getEnv("POLICY_MODEL", "./policies/model.conf"),
		PolicyFile:      getEnv("POLICY_FILE", "./policies/policy.csv"),
		OTelEndpoint:    getEnv("OTEL_EXPORTER_OTLP_ENDPOINT", ""),
		OTelServiceName: getEnv("OTEL_SERVICE_NAME", "bud2-backend"),
		OTelEnvironment: getEnv("OTEL_ENVIRONMENT", "development"),
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
