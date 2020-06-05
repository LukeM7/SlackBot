// Luke Manzitto
package main

import (
	"fmt"

	"github.com/slack-go/slack"
)

func main() {
	api := slack.New("xoxb-1117499265159-1151646799927-KJXqU3blMJaRjeIlgYWEPe7S")
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for message := range rtm.IncomingEvents {
		switch event := message.Data.(type) {
		case *slack.MessageEvent:
			go handleMessage(event)
		}
	}

}
func handleMessage(event *slack.MessageEvent) {
	fmt.Printf("%v\n", event)
}

/*
func response(rtm *slack.RTM, message *slack.MessageEvent, prefix string) {

	messageText := message.Text
	messageText = strings.TrimPrefix(messageText, prefix)
	messageText = strings.TrimSpace(messageText)
	messageText = strings.ToLower(messageText)

	gameStarts := map[string]bool{
		"rock paper scissors": true,
	}

	playGame := map[string]bool{
		"rock":     true,
		"paper":    true,
		"scissors": true,
	}

	textResponses := map[string]bool{
		"hello":      true,
		"what's up?": true,
	}

	help := map[string]bool{"help": true}

	if gameStarts[messageText] {
		rtm.SendMessage(rtm.NewOutgoingMessage("Choose rock, paper, or scissors", message.Channel))
	} else if playGame[messageText] {
		rtm.SendMessage(rtm.NewOutgoingMessage(rockPaperScissors(), message.Channel))
	} else if help[messageText] {
		rtm.SendMessage(rtm.NewOutgoingMessage(
			"Commands: \n\t@Gobot rock paper scissors *choice*\n\t @Gobot hello/what's up", message.Channel))
	} else if textResponses[messageText] {
		rtm.SendMessage(rtm.NewOutgoingMessage("Hey how are you doing?", message.Channel))
	}
}

func rockPaperScissors() string {
	gameResponse := map[int]string{
		0: "rock",
		1: "paper",
		2: "scissors",
	}
	randomIndex := rand.Int() % 3
	return gameResponse[randomIndex]
}
*/
