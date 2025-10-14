package bot

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	// "encoding/json"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/tidwall/gjson"
)

var (
	WeatherApiKey string
	BotToken string
	DbUri string
)

type User struct {
	UserId     string `bson:"user_id"`
	UserName   string `bson:"user_name"`
	GlobalName string `bson:"global_name"`
}

type Msg struct {
	MsgContent string `bson:"msg_content"`
}

type UserMsg struct {
	User User `bson:"user"`
	Msg  Msg  `bson:"msg"`
}

type Status struct {
	Status string
	Content string
}

func registerCommands(s *discordgo.Session){
	commands := []*discordgo.ApplicationCommand{
		{
			Name: "ping",
			Description: "Test apakah bot aktif",
		},
	}

	for _, cmd := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", cmd)
		if err != nil {
			fmt.Printf("Tidak bisa membuat command %v: %v\n", cmd.Name, err)
		}
	}
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

func Run() {
	fmt.Printf("bot token: %s\napi key: %s\n", BotToken, WeatherApiKey)
	
	// session init
	discord, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal("error in session initialization ", err)
	}

	// Event handler
	discord.AddHandler(newMsg)

	// Open session
	ok := discord.Open()
	if ok != nil {
		log.Fatal("error when trying to open session")
	}
	defer discord.Close()

	// slash command

	// bot running
	fmt.Println("Bot is running....")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func newMsg(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore bot message
	if m.Author.ID == s.State.User.ID {
		return
	}

	// msg respond
	// fmt.Println(m.Content)
	userInput := UserMsg {
		User: User{
			UserId: m.Author.ID,
			UserName: m.Author.Username,
			GlobalName: m.Author.GlobalName,
		},
		Msg: Msg{
			MsgContent: m.Content,
		},
	}
	fmt.Println(userInput.Msg)

	switch {
	case strings.Contains(m.Content, "weather"):
		s.ChannelMessageSend(m.ChannelID, "Saya bisa bantu")
	case strings.Contains(m.Content, "bot"):
		s.ChannelMessageSend(m.ChannelID, "hi there!")
	case strings.Contains(m.Content, "bego"):
		s.ChannelMessageSend(m.ChannelID, "bangsat kau!!")
	case strings.Contains(m.Content, "system-test-db-write"):

		jsonData, marshalErr := json.Marshal(m)
		if marshalErr != nil {
			log.Fatal(marshalErr)
		}
		userInput := gjson.GetBytes(jsonData, "author")
		
		fmt.Println(userInput.Raw)
		userData := UserMsg {
			User: User{
				UserId: userInput.Get("id").String(),
				UserName: userInput.Get("username").String(),
				GlobalName: userInput.Get("global_name").String(),
			},

			Msg: Msg{
				MsgContent: "hello world",
			},
		}

		err := dbWrite(userData)
		if err != nil {
			log.Fatal(err)
		}

		s.ChannelMessageSend(m.ChannelID, m.Content)

	case strings.Contains(m.Content, "msg-info"):
		jsonData, err := json.MarshalIndent(m, "", "	")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(jsonData))
		s.ChannelMessageSend(m.ChannelID, string(jsonData))
	}
	
}

