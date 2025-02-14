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
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "channel-id",
					Description: "if the parameter is required, the bot will enter the user's current channel",
					Required:    false,
				},
			},
		},
		{
			Name:        "leave",
			Description: "leave voice channel",
		},
		{
			Name:        "play-test",
			Description: "testing the bot to play music",
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
