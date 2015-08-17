package main

import "math/rand"

var choices = []string{"rock", "paper", "scissors"}

func randomAnswer() string {
	return choices[rand.Intn(len(choices))]
}

func rps(player string, computer string) string {
	return player //temporary
}
