package main

import (
	"log"
	"os"

	"github.com/bensiauu/discord-bot/config"
	handler "github.com/bensiauu/discord-bot/handlers"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

var (
	cfgPath = "CONFIG_PATH"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env file")
	}
}

func main() {

	buf, err := os.ReadFile(os.Getenv(cfgPath))
	if err != nil {
		log.Fatal("failed to read config file", err)
	}

	cfg := &config.Config{}

	err = yaml.Unmarshal(buf, cfg)
	if err != nil {
		log.Fatal("could not parse config file")
	}
	bot, err := discordgo.New("Bot " + cfg.Authorization.Token)
	if err != nil {
		log.Fatal("failed to initialize bot")
	}

	// init handlers for bot
	// Register messageCreate func as callback for MessageCreate events
	handler := handler.NewHandler()
	handler.InitHandlers(bot)

	bot.Identify.Intents = discordgo.IntentsGuildMessages

}
