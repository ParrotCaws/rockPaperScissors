package main

import (
	"fmt"
	"math/rand"
	"time"
)

var choices = []string{"rock", "paper", "scissors"}
var p = fmt.Println

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
	var choice1 = randomAnswer()
	var choice2 = randomAnswer()
	var result = choiceCompare(choice1, choice2)
	p(result)
	p("player 1's choice: ", choice1)
	p("player 2's choice: ", choice2)
}
