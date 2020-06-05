// Luke Manzitto
package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/slack-go/slack"
)

func main() {
	api := slack.New("xoxb-1117499265159-1151646799927-KJXqU3blMJaRjeIlgYWEPe7S")
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for message := range rtm.IncomingEvents {
		switch event := message.Data.(type) {
		case *slack.MessageEvent:
			info := rtm.GetInfo()
			prefix := fmt.Sprintf("<@%s>", info.User.ID)
			go response(rtm, event, prefix)
		}
	}

}

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
		rtm.SendMessage(rtm.NewOutgoingMessage(rockPaperScissors(messageText), message.Channel))
	} else if help[messageText] {
		rtm.SendMessage(rtm.NewOutgoingMessage(
			"Commands: \n\t@Gobot rock paper scissors *choice*\n\t @Gobot hello/what's up", message.Channel))
	} else if textResponses[messageText] {
		rtm.SendMessage(rtm.NewOutgoingMessage("Hey how are you doing?", message.Channel))
	}
}

func rockPaperScissors(userResponse string) string {
	gameResponse := map[int]string{
		0: "rock",
		1: "paper",
		2: "scissors",
	}
	randomIndex := rand.Intn(3)
	fmt.Println(randomIndex)
	fmt.Println("I chose " + gameResponse[randomIndex])

	if gameResponse[randomIndex] == userResponse {
		return "Draw"
	} else if (userResponse == "scissors" && gameResponse[randomIndex] == "paper") || (userResponse == "paper" && gameResponse[randomIndex] == "rock") || (userResponse == "rock" && gameResponse[randomIndex] == "scissors") {
		return "You Win!"
	} else {
		return "I WIN!"
	}
}
