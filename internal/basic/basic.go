package basic

import (
	"fmt"
	"ishikawayae/internal/common"

	"github.com/bwmarrin/discordgo"
)

func Ping(s *discordgo.Session, i *discordgo.InteractionCreate, c *common.Config) {
	delay := s.HeartbeatLatency()
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Delay time: %v", delay),
		},
	})
}
