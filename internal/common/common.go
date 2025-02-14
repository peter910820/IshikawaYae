package common

import (
	"github.com/bwmarrin/discordgo"
)

type Config struct {
	Bot         *discordgo.Session
	Token       string
	AppID       string
	StartStatus string
	VcMap       map[string]*discordgo.VoiceConnection
}
