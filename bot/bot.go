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
	"github.com/tidwall/gjson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	WeatherApiKey string
	BotToken string
	DbUri string
	
	defaultMemberPermissions int64 = discordgo.PermissionManageGuild

	// command list
	commands = []*discordgo.ApplicationCommand{

		{
			Name: "basic-command",
			Description: "Test the basic functionality of slash command",
		},

		{
			Name: "server-config",
			Description: "Cmmand for demonstration of default command permisssion",
			DefaultMemberPermissions: &defaultMemberPermissions,
		},

		{
			Name: "options",
			Description: "for testing options capability in slash commands",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type: discordgo.ApplicationCommandOptionString,
					Name: "string-option",
					Description: "string-option",
					Required: true,
				},
				
				{
					Type: discordgo.ApplicationCommandOptionInteger,
					Name: "integer-option",
				},
			},
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"basic-command": func(s *discordgo.Session, i *discordgo.InteractionCreate){
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Hello this is basic slash command",
				},
			})
		},

		"server-config": func(s *discordgo.Session, i *discordgo.InteractionCreate){
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Server status: all good \nBot status: all good",
				},
			})
		},

		"options": func(s *discordgo.Session, i *discordgo.InteractionCreate){
			
			options :=  i.ApplicationCommandData().Options

			optionsMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionsMap[opt.Name] = opt
			}
			
			msgformat := "You learned how to use command options! " +
				"Take a look at the value(s) you entered:\n"
			margs := make([]interface{}, 0, len(options))

			if opt, ok := optionsMap["string-option"]; ok {
				margs = append(margs, opt.StringValue())
				msgformat += "> string-option: %s\n"
			}
			
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf(
						msgformat,
						margs...
					),
				},
			})
		},
	}
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

func Main() {
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

