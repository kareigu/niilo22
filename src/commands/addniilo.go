package commands

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type checkIDPost struct {
	Id        string `json:"id"`
	Permlevel int32  `json:"permlevel"`
}

type checkIDRes struct {
	Id        string
	Permlevel int32
	Permitted bool
}

type addNiiloRes struct {
	Message string
}

var addNiiloInfo = discordgo.ApplicationCommand{
	Name:        "addniilo",
	Description: "Add a new niilo quote. Only accessible to VIP+",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "quote",
			Description: "Enter the new quote to be added",
			Required:    true,
		},
	},
}

func addNiiloCmd(s *discordgo.Session, i *discordgo.InteractionCreate) {
	userID := "empty"
	if i.Member != nil {
		userID = i.Member.User.ID
	} else if i.User != nil {
		userID = i.User.ID
	}

	reqBody := checkIDPost{
		Id:        userID,
		Permlevel: 3,
	}

	res := new(checkIDRes)

	err := postData("https://misi.karei.dev/api/v1/tools/checkid", &res, &reqBody)
	if err != nil {
		log.Print(err)
	}

	if res.Permitted && res.Permlevel >= 3 {
		key := os.Getenv("KEY")
		new_quote := strings.ReplaceAll(i.Data.Options[0].StringValue(), " ", "%20")
		add_url := fmt.Sprintf("https://misi.karei.dev/api/v1/niilo/add?secret=%s&text=%s",
			key,
			new_quote)

		add_res, err := getDataRaw(add_url)
		if err != nil {
			log.Print(err)
			add_res = "Hyvin menee, mut ei mee (Virhe tapahtui)"
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionApplicationCommandResponseData{
				Content: add_res,
			},
		})
	} else {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionApplicationCommandResponseData{
				Content: "Mitä vittua tää ny meinaa? (VIP+ vaadittu)",
			},
		})
	}
}
