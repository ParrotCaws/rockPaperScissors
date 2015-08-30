package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	choices                      = []string{"rock", "paper", "scissors"}
	put                          = fmt.Println
	get                          = fmt.Scanf
	running                      bool
	player1choice, player2choice string
)

func randomAnswer() string {
	return choices[rand.Intn(len(choices))]
}

func isValid(choice string) bool {
	for _, c := range choices {
		if choice == c {
			return true
		}
	}
	return false
}

func choiceCompare(choice1, choice2 string) string {
	switch {
	case !isValid(choice1):
		return "Player one's choice was not valid!"
	case !isValid(choice2):
		return "Player two's choice was not valid!"
	case choice1 == choice2:
		return "There was a tie!"
	case (choice1 == "rock" && choice2 == "scissors") || (choice1 == "paper" && choice2 == "rock") || (choice1 == "scissors" && choice2 == "paper"):
		return "Player one wins!"
	default:
		return "Player two wins!"
	}
}
func playAgain() bool {
	put("Would you like to play again? (y/n)")

	var again string
	get("%s", &again)
	again = strings.ToLower(again)

	if again == "y" {
		return true
	} else if again == "n" {
		return false
	} else {
		put("That's an invalid response!")
		return playAgain()
	}
}
func game() {
	for running == true {
		put("Would you like to play with 0, 1, or 2 players?")

		var players string
		get("%s", &players)
		if players == "1" || players == "2" {
			put("What is your choice? (rock, paper, or scissors)")
			get("%s", &player1choice)
		} else {
			put("The player 1 computer is choosing...")
			player1choice = randomAnswer()
		}
		if players == "2" {
			put("what is your opponents choice?")
			get("%s", &player2choice)
		} else {
			put("The player 2 computer is choosing...")
			player2choice = randomAnswer()
		}
		player1choice = strings.ToLower(player1choice)
		player2choice = strings.ToLower(player2choice)

		put(choiceCompare(player1choice, player2choice))
		put("Player 1 chose ", player1choice)
		put("Player 2 chose ", player2choice)

		running = playAgain()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	running = true
	put("Welcome to Rock Paper Scissors!")
	game()
	put("Thanks for playing Rock Paper Scissors, by ParrotCaws")
	put("Fork me on github! github.com/ParrotCaws")
}
