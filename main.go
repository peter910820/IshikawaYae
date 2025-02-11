package main

import (
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"ishikawayae/internal/bot"
)

var (
	c bot.Config
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

	c = bot.Config{
		Token:       os.Getenv("TOKEN"),
		AppID:       os.Getenv("APP_ID"),
		StartStatus: os.Getenv("Start_Status"),
	}

	c.Bot, err = discordgo.New("Bot " + c.Token)
	if err != nil {
		logrus.Fatal(err)
	}

	c.Bot.AddHandler(ready)

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
}
