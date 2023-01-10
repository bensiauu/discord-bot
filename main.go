package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

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

	err = bot.Open()
	if err != nil {
		log.Fatal("failed to open connection", err)
	}

	fmt.Println("Bot is now running. press CTRL-C to exit.")
	// make new channel called sc to signify that we want to listen to OS signals
	sc := make(chan os.Signal, 1)
	// notify signal for all signals we want to listen to, namely: SIGINT, SIGTERM and common system interrupt and kill signals
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Kill, os.Interrupt)
	// halts program from executing the next line until one of those signals are received
	<-sc

	bot.Close()

}
