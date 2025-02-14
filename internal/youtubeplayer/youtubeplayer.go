package youtubeplayer

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"

	"ishikawayae/internal/common"
)

func Join(s *discordgo.Session, i *discordgo.InteractionCreate, c *common.Config) {
	data, ok := i.Interaction.Data.(discordgo.ApplicationCommandInteractionData)
	if !ok {
		logrus.Warnln("join command assertion failed")
		return
	}
	var channelID string
	for _, opt := range data.Options {
		if opt.Name == "channel-id" {
			channelID = opt.Value.(string)
			break
		}
	}

	vc, err := c.Bot.ChannelVoiceJoin(i.GuildID, channelID, false, false)
	if err != nil {
		logrus.Fatalf("join: %s failed: %s", channelID, err)
		return
	}
	c.VcMap[i.GuildID] = vc
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "加入語音頻道成功",
		},
	})
}

func Leave(s *discordgo.Session, i *discordgo.InteractionCreate, c *common.Config) {
	vc, exists := c.VcMap[i.GuildID]
	if !exists || vc == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "機器人未加入語音頻道",
			},
		})
		return
	}

	vc.Disconnect()
	delete(c.VcMap, i.GuildID)
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "已退出語音頻道",
		},
	})

}

func PlayTest(s *discordgo.Session, i *discordgo.InteractionCreate, c *common.Config) {
	// check voiceConnection
	var voiceConnection *discordgo.VoiceConnection
	for _, v := range c.Bot.VoiceConnections {
		if v.GuildID == i.GuildID {
			voiceConnection = v
			break
		}
	}

	voiceConnection.Speaking(true)
}
