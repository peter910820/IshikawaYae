package main

import (
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"ishikawayae/internal/bot"
	"ishikawayae/internal/common"
)

var (
	c common.Config
)

func main() {
	// logrus settings
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.DebugLevel)
	// load .env
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatal(err)
	}

	c = common.Config{
		Token:       os.Getenv("TOKEN"),
		AppID:       os.Getenv("APP_ID"),
		StartStatus: os.Getenv("Start_Status"),
		VcMap:       make(map[string]*discordgo.VoiceConnection),
	}

	c.Bot, err = discordgo.New("Bot " + c.Token)
	if err != nil {
		logrus.Fatal(err)
	}

	c.Bot.AddHandler(ready)
	c.Bot.AddHandler(onInteraction)

	err = c.Bot.Open()
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("bot is now running. Press CTRL+C to exit.")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	interruptSignal := <-ch
	c.Bot.Close()
	logrus.Info(interruptSignal)
}

func ready(s *discordgo.Session, r *discordgo.Ready) {
	s.UpdateGameStatus(0, c.StartStatus)
	bot.CleanCommand(c.Bot)
	bot.RegisterCommand(s)
}

func onInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	go bot.HandleInteraction(s, i, &c)
}
