package handler

import "github.com/bwmarrin/discordgo"

type Handler interface {
	InitHandlers(bot *discordgo.Session)
	MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate)
}

type handler struct {
}

func (h *handler) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

}

func (h *handler) InitHandlers(bot *discordgo.Session) {
	bot.AddHandler(h.MessageCreate)
}

func NewHandler() Handler {
	return &handler{}
}
