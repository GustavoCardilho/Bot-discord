package commands

import "github.com/bwmarrin/discordgo"

var CommandLabels = []*discordgo.ApplicationCommand{
	{
		Name:        "basic-command",
		Description: "Basic command",
	},
}
