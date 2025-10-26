package handlers

import (
	"errors"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

const (
	MethodWrite = "write"
	MethodRead = "read"
)

func SlashCommandHandler (s *discordgo.Session, i *discordgo.InteractionCreate)  {
	if i.Type != discordgo.InteractionApplicationCommand {
		fmt.Printf("erorr in initiating slash command")
	}

	commandName := i.ApplicationCommandData().Name
	switch commandName {
	case "jadwal" :
		err := jadwalKuliah(s, i)
		if err != nil {
			fmt.Printf("jadwal slash command failed")
		}
	}
}

func jadwalKuliah (s *discordgo.Session, i *discordgo.InteractionCreate) error {
	// accesing database
	
	
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "jadwal kuliah testing",
		},
	})

	if err != nil {
		return errors.New("error in /jadwalkuliah command respond")
	}
	return nil
}



