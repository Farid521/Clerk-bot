package main

import (
	"github.com/joho/godotenv"
	"clerk-bot/bot"
	"log"
	"os"
)

func main() {
	// load the .env data
	err := godotenv.Load()

	if err != nil {
		log.Fatal("env file cannot be loaded")
	}

	BotToken, found := os.LookupEnv("BOT_TOKEN")
	if !found {
		log.Fatal("bot token cannot be found")
	}
	WeatherApiKey, found := os.LookupEnv("WEATHER_API_KEY")
	if !found {
		log.Fatal("weather api key cannot be found")
	}
	DbUri, found := os.LookupEnv("DB_URI")
	if !found {
		log.Fatal("DB Uri not found")
	}

	// bot initialization
	bot.BotToken = BotToken
	bot.WeatherApiKey = WeatherApiKey
	bot.DbUri = DbUri
	bot.Run()
}
