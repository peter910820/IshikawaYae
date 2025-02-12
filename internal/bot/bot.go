package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Bot         *discordgo.Session
	Token       string
	AppID       string
	StartStatus string
}

func RegisterCommand(s *discordgo.Session) {
	BasicCommands(s)
	YoutubePlayerCommands(s)
}

func CleanCommand(bot *discordgo.Session) {
	commands, err := bot.ApplicationCommands(bot.State.User.ID, "")
	if err != nil {
		logrus.Warnln(err)
	}

	for _, cmd := range commands {
		err := bot.ApplicationCommandDelete(bot.State.User.ID, "", cmd.ID)
		if err != nil {
			logrus.Warnf("delete %s failed: %s.\n", cmd.Name, err)
		} else {
			logrus.Infof("delete %s success.\n", cmd.Name)
		}
	}
}
