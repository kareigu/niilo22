package commands

import (
	"fmt"
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
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "search",
			Description: "Type in a search term or the quote number to retrieve that quote. $ gets the newest quote",
			Required:    false,
		},
	},
}

func niiloCmd(s *discordgo.Session, i *discordgo.InteractionCreate) {
	quote := new(niiloRes)
	search_param := ""
	if len(i.Data.Options) > 0 {
		search_param = i.Data.Options[0].StringValue()
	}

	url := fmt.Sprintf("https://misi.karei.dev/api/v1/niilo?search=%s", search_param)
	text := ""
	err := getData(url, &quote)
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
