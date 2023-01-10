package handler

import "github.com/bwmarrin/discordgo"

type Handler interface {
	InitHandlers(bot *discordgo.Session)
	messageCreate(s *discordgo.Session, m *discordgo.MessageCreate)
}

type handler struct {
}

func (h *handler) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "pong")
	}

	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "ping")
	}
}

func (h *handler) InitHandlers(bot *discordgo.Session) {
	bot.Identify.Intents = discordgo.IntentsGuildMessages

	bot.AddHandler(h.messageCreate)
}

func NewHandler() Handler {
	return &handler{}
}
