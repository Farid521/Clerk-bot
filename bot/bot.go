package bot

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	// "encoding/json"

	"github.com/bwmarrin/discordgo"
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

func Main() {
    // session init
    discord, err := discordgo.New("Bot " + BotToken)
    if err != nil {
        log.Fatal("error in session initialization ", err)
    }

    // Open session
    ok := discord.Open()
    if ok != nil {
        log.Fatal("error when trying to open session")
    }
    defer discord.Close()

	// Event message handler
    discord.AddHandler(newMsg)

	// commands initilization
	commandsList := make([]*discordgo.ApplicationCommand, 0, len(Commands))
	for key := range Commands {
		v := Commands[key]
		vpointer := &v
		commandsList = append(commandsList, vpointer)
	}

	for _, command := range commandsList {
		_, err := discord.ApplicationCommandCreate(discord.State.User.ID, "", command)
		if err != nil {
			fmt.Printf("error registering command")
		} else {
			fmt.Printf("Succesfully registered %s command\n------------\n", command.Name)
		}
	}

	// slash command handler
	discord.AddHandler(SlashCommandHandler)

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

