package config

import "os"

type Config struct {
	DBUrl      string
	RedisAddr  string
	ServerPort string
}

func Load() *Config {
	return &Config{
		DBUrl:      getEnv("DB_URL", "user:pass@tcp(localhost:3306)/todo_db?charset=utf8mb4&parseTime=True&loc=Local"),
		RedisAddr:  getEnv("REDIS_ADDR", "localhost:6379"),
		ServerPort: getEnv("SERVER_PORT", ":8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
