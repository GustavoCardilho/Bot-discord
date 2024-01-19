package boostrap

import (
	"github.com/GustavoCardilho/Bot-discord/infra/commands"
	"github.com/GustavoCardilho/Bot-discord/infra/handlers"
	"github.com/bwmarrin/discordgo"
)

func addingHandlers() {
	session.AddHandler(handlers.MessageCreate)
	session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commands.Handlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}
