package commands

import (
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
