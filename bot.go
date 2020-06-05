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
			go response(rtm, event)
		case *slack.RTMError:
			fmt.Println(event.Error())
		case *slack.InvalidAuthEvent:
			fmt.Println("Bad authentication token")
			return

		}
	}

}

func response(rtm *slack.RTM, message *slack.MessageEvent) {

	messageText := message.Text
	messageText = strings.TrimSpace(messageText)
	messageText = strings.ToLower(messageText)

	// Initiates Rock,Paper,Scissors Game
	gameStarts := map[string]bool{
		"rock paper scissors": true,
	}

	// If user inputs one of these the bot will pick and return the outcome
	playGame := map[string]bool{
		"rock":     true,
		"paper":    true,
		"scissors": true,
	}
	// Responses to common greetings
	textResponses := map[string]bool{
		"hi":         true,
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
			"Commands: \n\tTo play a game: @Gobot rock paper scissors *choice*\n\tTo say hi: @Gobot hello/what's up", message.Channel))
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
