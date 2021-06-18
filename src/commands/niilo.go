package commands

import (
	"log"

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

func niiloCmd(s *discordgo.Session, i *discordgo.InteractionCreate) {
	quote := new(niiloRes)
	text := ""
	err := getData("https://misi.mxrr.dev/api/v1/niilo", &quote)
	if err != nil {
		text = "Vittu kun ei tää yhdellä kädellä onnistu"
		log.Printf("Error getting niilo data: %v", err)
	} else {
		text = quote.Text
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionApplicationCommandResponseData{
			Content: text,
		},
	})
}
