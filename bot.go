// Luke Manzitto
/*
* Sources used to Learn:
* https://rsmitty.github.io/Slack-Bot/
* https://x-team.com/blog/writing-slackbots-with-goroutines/
* https://www.youtube.com/watch?v=zkB_c3cgtd0
 */
package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/slack-go/slack"
)

// Global game check var
var gameCheck bool = false

func main() {
	// Should use environment variable for good practice
	api := slack.New("xoxb-1117499265159-1151646799927-KJXqU3blMJaRjeIlgYWEPe7S")
	rtm := api.NewRTM()
	go rtm.ManageConnection() // Goroutine manages websocket connections
	for message := range rtm.IncomingEvents {
		switch event := message.Data.(type) {
		// If a message is recieved
		case *slack.MessageEvent:
			// Used https://rsmitty.github.io/Slack-Bot/ 's method to find prefix
			info := rtm.GetInfo()
			prefix := fmt.Sprintf("<@%s>", info.User.ID)
			response(rtm, event, prefix)
		// If the RTM gets an error
		case *slack.RTMError:
			fmt.Println(event.Error())
		// If the authentaction token doesn't work
		case *slack.InvalidAuthEvent:
			fmt.Println("Bad authentication token")
			return
		}
	}
}

// Bot responses to user input
func response(rtm *slack.RTM, message *slack.MessageEvent, prefix string) {

	// puts user message in all lowercase so we don't have to account for variations in upper/lowercase
	messageText := message.Text
	messageText = strings.TrimPrefix(messageText, prefix)
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
		gameCheck = true
		rtm.SendMessage(rtm.NewOutgoingMessage("Choose rock, paper, or scissors", message.Channel))
	} else if playGame[messageText] && gameCheck {
		gameCheck = false
		rtm.SendMessage(rtm.NewOutgoingMessage(rockPaperScissors(messageText, rtm, message), message.Channel))
	} else if playGame[messageText] && !gameCheck {
		rtm.SendMessage(rtm.NewOutgoingMessage("Please initiate game", message.Channel))
	} else if help[messageText] {
		rtm.SendMessage(rtm.NewOutgoingMessage(
			"Commands: \n\tTo play a game: @Gobot rock paper scissors\n\tTo say hi: @Gobot hi/hello/what's up", message.Channel))
	} else if textResponses[messageText] {
		rtm.SendMessage(rtm.NewOutgoingMessage("Hello there", message.Channel))
	}
}

// Rock, paper, scissors logic
func rockPaperScissors(userResponse string, rtm *slack.RTM, message *slack.MessageEvent) string {
	gameResponse := map[int]string{
		0: "rock",
		1: "paper",
		2: "scissors",
	}
	randomIndex := rand.Intn(3)
	fmt.Println(randomIndex)
	fmt.Println("I chose " + gameResponse[randomIndex])
	rtm.SendMessage(rtm.NewOutgoingMessage("I chose "+gameResponse[randomIndex], message.Channel))
	if gameResponse[randomIndex] == userResponse {
		return "DRAW"
	} else if (userResponse == "scissors" && gameResponse[randomIndex] == "paper") || (userResponse == "paper" && gameResponse[randomIndex] == "rock") || (userResponse == "rock" && gameResponse[randomIndex] == "scissors") {
		return "YOU WIN!"
	} else {
		return "I WIN!"
	}
}
