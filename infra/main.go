package main

import (
	"log"
	"os"

	"github.com/GustavoCardilho/Bot-discord/infra/boostrap"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	integerOptionMinValue          = 1.0
	dmPermission                   = false
	defaultMemberPermissions int64 = discordgo.PermissionManageServer
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}
	token := os.Getenv("TOKEN")
	boostrap.Run(token)
}
