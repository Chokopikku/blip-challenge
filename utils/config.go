package utils

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	LogFilePath     string
	RankingFilePath string
	WeightCommits   float64
	WeightFiles     float64
	WeightLines     float64
}

func LoadConfig() *Config {
	_ = godotenv.Load(".env")

	return &Config{
		LogFilePath:     getEnv("LOG_FILE", "app.log"),
		RankingFilePath: getEnv("RANKING_FILE", "ranking.json"),
		WeightCommits:   getEnvAsFloat("WEIGHT_COMMITS", 1.0),
		WeightFiles:     getEnvAsFloat("WEIGHT_FILES", 0.4),
		WeightLines:     getEnvAsFloat("WEIGHT_LINES", 0.2),
	}
}

func getEnv(key string, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

func getEnvAsFloat(key string, defaultVal float64) float64 {
	valStr := os.Getenv(key)
	if val, err := strconv.ParseFloat(valStr, 64); err == nil {
		return val
	}
	return defaultVal
}
