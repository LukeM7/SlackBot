// Luke Manzitto
package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/slack-go/slack"
)

func main() {

	api := slack.New("xoxp-1117499265159-1125508635446-1190043385088-54427fa1679c89e8725d604838afc783")
	rtm := api.NewRTM()
	go rtm.ManageConnection()
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Print("Event Received: ")
			switch ev := msg.Data.(type) {

			case *slack.MessageEvent:
				info := rtm.GetInfo()

				text := ev.Text
				text = strings.TrimSpace(text)
				text = strings.ToLower(text)

				matched, _ := regexp.MatchString("dark souls", text)

				if ev.User != info.User.ID && matched {
					rtm.SendMessage(rtm.NewOutgoingMessage("\\[T]/ Praise the Sun \\[T]/", ev.Channel))
				}

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				return

			default:
				// Take no action
			}
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

