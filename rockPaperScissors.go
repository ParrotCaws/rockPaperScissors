package main

import (
	"fmt"
	"math/rand"
	"time"
)

var choices = []string{"rock", "paper", "scissors"}
var put = fmt.Println
var get = fmt.Scanf
var running bool
var player1choice string
var player2choice string

func randomAnswer() string {
	return choices[rand.Intn(len(choices))]
}

func choiceCompare(choice1 string, choice2 string) string {
	switch {
	case choice1 == "rock":
		if choice2 == "rock" {
			return "It was a tie!!"
		} else if choice2 == "paper" {
			return "Player 2 wins!"
		} else {
			return "Player 1 wins!"
		}
	case choice1 == "paper":
		if choice2 == "rock" {
			return "Player 1 wins!"
		} else if choice2 == "paper" {
			return "It was a tie!"
		} else {
			return "Player 2 wins!"
		}
	case choice1 == "scissors":
		if choice2 == "rock" {
			return "Player 2 wins!"
		} else if choice2 == "paper" {
			return "Player 1 wins!"
		} else {
			return "It was a tie!"
		}
	default:
		return "There was an error!"
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	running = true

	for running == true {
		put("Welcome to Rock Paper Scissors!")
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

		put(choiceCompare(player1choice, player2choice))
		put("Player 1 chose ", player1choice)
		put("Player 2 chose ", player2choice)

		put("Would you like to play again? (y/n)")
		var playAgain string
		get("%s", &playAgain)

		if playAgain == "y" {
			running = true
		} else {
			running = false
		}
	}
	put("Thanks for playing Rock Paper Scissors, by ParrotCaws")
	put("Fork me on github! github.com/ParrotCaws")
}
