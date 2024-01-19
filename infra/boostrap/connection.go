package boostrap

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func ConnectionBot() {
	var err error
	token := os.Getenv("TOKEN")
	session, err = discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Error creating Discord session: ", err)
	}
}
