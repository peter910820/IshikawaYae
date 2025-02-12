package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func YoutubePlayerCommands(s *discordgo.Session) {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "join",
			Description: "join into voice channel",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "voice-cahnnle-id",
					Description: "if the parameter is required, the bot will enter the user's current channel",
					Required:    false,
				},
			},
		},
	}
	for _, cmd := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", cmd)
		if err != nil {
			logrus.Error(err)
			return
		}
	}
	logrus.Infof("add command modules %s success.\n", "youtubeplayer")
}
