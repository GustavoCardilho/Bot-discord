package boostrap

import (
	"log"
	"os"

	"github.com/GustavoCardilho/Bot-discord/infra/commands"
	"github.com/bwmarrin/discordgo"
)

func RegisterCommands() {
	GuildID := os.Getenv("GUILD_ID")
	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands.CommandsDetails))
	for i, v := range commands.CommandsDetails {
		cmd, err := session.ApplicationCommandCreate(session.State.User.ID, GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}
}
