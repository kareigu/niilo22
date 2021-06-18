package commands

import (
	"github.com/bwmarrin/discordgo"
)

var niiloInfo = discordgo.ApplicationCommand{
	Name:        "niilo",
	Description: "Post a niilo quote",
}

func niiloCmd(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionApplicationCommandResponseData{
			Content: "onh",
		},
	})
}
