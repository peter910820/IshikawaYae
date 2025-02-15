package youtubeplayer

import (
	"log"
	"os/exec"
	"time"

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

	cmd := exec.Command("ffmpeg", "-i", "filpath", "-f", "opus", "-ar", "48000", "-ac", "2", "-b:a", "64k", "-loglevel", "debug", "-")
	// 透過管道接收 ffmpeg 的輸出
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal("無法建立管道:", err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal("啟動 ffmpeg 錯誤:", err)
	}

	// 定義緩衝區
	buf := make([]byte, 1024) // 每個 Opus 音頻框架的大小

	// 發送音頻流到 Discord 語音頻道
	for {
		n, err := stdout.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Println("讀取音頻資料時發生錯誤:", err)
			break
		}

		// 將音頻資料寫入 OpusSend channel
		voiceConnection.OpusSend <- buf[:n]

		// 控制播放速率，這裡加入延遲，確保每次送出資料的頻率正確
		time.Sleep(20 * time.Millisecond)
	}

	if err := cmd.Wait(); err != nil {
		log.Println("ffmpeg 命令錯誤:", err)
	}
}
