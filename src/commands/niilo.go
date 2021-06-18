package commands

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

type niiloRes struct {
	Id     string `json:"_id"`
	Text   string
	Number uint64
}

var niiloInfo = discordgo.ApplicationCommand{
	Name:        "niilo",
	Description: "Post a niilo quote",
}

var client = &http.Client{Timeout: 10 * time.Second}

func niiloCmd(s *discordgo.Session, i *discordgo.InteractionCreate) {
	res, err := client.Get("https://misi.mxrr.dev/api/v1/niilo")
	if err != nil {
		log.Print(err)
	}

	defer res.Body.Close()

	quote := new(niiloRes)
	json.NewDecoder(res.Body).Decode(&quote)
	log.Print(quote)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionApplicationCommandResponseData{
			Content: "onh",
		},
	})
}
