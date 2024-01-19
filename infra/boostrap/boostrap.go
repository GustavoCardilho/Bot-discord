package boostrap

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var session *discordgo.Session

func Run(token string) {
	ConnectionBot()
	session.Identify.Intents = discordgo.IntentsGuildMessages
	errOpenSession := session.Open()
	if errOpenSession != nil {
		fmt.Println("error opening connection,", errOpenSession)
		return
	}
	RegisterCommands()
	addingHandlers()
	Exit()
}
