package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"

	"github.com/nlopes/slack"
	"github.com/spf13/viper"
)

var (
	SlackAPIToken = "" // to be loaded from file
	yodaisms      []string
	Me            = "" // eventually the bot's name
)

func init() {

	// Get Slack API Token

	// os.Getenv("SLACK_BOT_TOKEN")

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Config error: %s \n", err))
	} else {
		SlackAPIToken = viper.GetString("slack_api_token")
	}

	// Load data file

	file, err := ioutil.ReadFile("./yoda_said.json")
	if err != nil {
		panic(fmt.Errorf("Data file error: %s \n", err))
	} else {
		err := json.Unmarshal(file, &yodaisms)
		if err != nil {
			panic(fmt.Errorf("Json parse error: %s \n", err))
		}
	}

	// Seed rand
	rand.Seed(time.Now().UTC().UnixNano())

}

// TODO list:
// [] don't hardcode the channel

func main() {
	api := slack.New(SlackAPIToken)

	//api.SetDebug(true)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:

			switch event := msg.Data.(type) {
			case *slack.ConnectedEvent:
				fmt.Println("Info: ", event.Info)
				response := "Beep Boop. I am a bot. I listen for the keyword `WWYS?` (i.e., What Would Yoda Say?)"
				rtm.SendMessage(rtm.NewOutgoingMessage(response, "C3MMS8G6B"))
			case *slack.MessageEvent:
				if event.Type == "message" && strings.Contains(event.Text, "WWYS?") {
					rtm.SendMessage(rtm.NewOutgoingMessage(yodaisms[rand.Intn(len(yodaisms))], event.Channel))
				}
			default:
				//empty case
			}

		}
	}

}
