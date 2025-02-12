package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"

	"ishikawayae/internal/basic"
	"ishikawayae/internal/youtubeplayer"
)

var commandHandlerMap = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"ping": basic.Ping,

	"join": youtubeplayer.Join,
}

func HandleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	cmdName := i.ApplicationCommandData().Name
	if handler, exists := commandHandlerMap[cmdName]; exists {
		handler(s, i)
	} else {
		logrus.Warnf("command not found: %s", cmdName)
	}
}
