package main

import (
	"clerk-bot/src"
	"clerk-bot/config"
)

func main() {
	// bot initialization
	config.LoadEnv()
	src.Main()
}
