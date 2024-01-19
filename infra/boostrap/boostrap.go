package boostrap

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/GustavoCardilho/Bot-discord/infra/commands"
	"github.com/GustavoCardilho/Bot-discord/infra/handlers"
	"github.com/bwmarrin/discordgo"
)

var session *discordgo.Session

func Run(token string) {
	var err error
	session, err = discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Error creating Discord session: ", err)
	}
	session.AddHandler(handlers.MessageCreate)
	session.Identify.Intents = discordgo.IntentsGuildMessages
	err = session.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

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

	session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commands.Handlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	session.Close()
}
