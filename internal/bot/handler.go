package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"

	"ishikawayae/internal/basic"
	"ishikawayae/internal/common"
	"ishikawayae/internal/youtubeplayer"
)

var commandHandlerMap = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate, c *common.Config){
	"ping": basic.Ping,

	"join":  youtubeplayer.Join,
	"leave": youtubeplayer.Leave,
}

func HandleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate, c *common.Config) {
	cmdName := i.ApplicationCommandData().Name
	if handler, exists := commandHandlerMap[cmdName]; exists {
		handler(s, i, c)
	} else {
		logrus.Warnf("command not found: %s", cmdName)
	}
}
