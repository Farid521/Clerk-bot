package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var (
	BotToken string
	WeatherApiKey string
	DbUri string
)

func LoadEnv() {
	err := godotenv.Load(filepath.Join("../", ".env"))
	if err != nil {
		log.Fatal("env file cannot be loaded	",err)
	}

	BotToken = mustGetEnv("BOT_TOKEN")
	WeatherApiKey = mustGetEnv("WEATHER_API_KEY")
	DbUri = mustGetEnv("DB_URI")
}

func mustGetEnv(key string) string {
	value, found := os.LookupEnv(key)
	if !found {
		log.Fatalf("Environment variable %s not found", key)
	}
	return value
}