package commands

import "github.com/bwmarrin/discordgo"

var CommandsDetails = []*discordgo.ApplicationCommand{
	{
		Name:        "basic-command",
		Description: "Basic command",
	},
}

var Handlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"basic-command": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Hey there! Congratulations, you just executed your first slash command",
			},
		})
	},
}
