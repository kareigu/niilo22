package commands

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	List = []*discordgo.ApplicationCommand{
		&niiloInfo,
	}

	Handlers = map[string]func(
		s *discordgo.Session,
		i *discordgo.InteractionCreate){
		niiloInfo.Name: niiloCmd,
	}
)

var client = &http.Client{Timeout: 10 * time.Second}

func getData(url string, target interface{}) error {
	res, err := client.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(&target)
}
