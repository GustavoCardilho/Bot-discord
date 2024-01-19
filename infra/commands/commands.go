package commands

import (
	CommandExecute "github.com/GustavoCardilho/Bot-discord/infra/commands/execute"
	"github.com/bwmarrin/discordgo"
)

var CommandsDetails = CommandLabels

var Handlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"basic-command": CommandExecute.BasicCommandExecute,
}
