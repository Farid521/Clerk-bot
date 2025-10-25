package bot

import (
	"errors"
	"fmt"
	"context"
	"time"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SlashCommandHandler (s *discordgo.Session, i *discordgo.InteractionCreate) error {
	if i.Type != discordgo.InteractionApplicationCommand {
		return errors.New("error in initiating slash command handler")
	}

	switch i.ApplicationCommandData().Name {
	case "jadwal" :
		err := jadwalKuliah(s, i)
		if err != nil {
			fmt.Printf("jadwal slash command failed")
		}
	}

	return nil
}

func jadwalKuliah (s *discordgo.Session, i *discordgo.InteractionCreate) error {
	
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


func dbWrite(content UserMsg) (error) {
	clientOptions := options.Client().ApplyURI(DbUri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err !=  nil {
		return err
	}

	fmt.Println("connected to mongoDB Atlas")

	// input
	collection := client.Database("DiscordBot").Collection("TestData")
	userInput := content
	result, err := collection.InsertOne(ctx, userInput)
	if err != nil {
		fmt.Println("error in inserting data")
		return err
	}

	fmt.Println("succesfully inserting data")
	fmt.Println(result.InsertedID)
	return nil
}


