package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	outcomeMap := map[selection]map[selection]int{
		// opponent selection: [myselection: myscore]
		Rock: {Paper: 6, Rock: 3, Scissors: 0},
		Paper: {Scissors: 6, Paper: 3, Rock: 0},
		Scissors: {Rock: 6, Scissors: 3, Paper: 0},
	}
	totalScore := 0
	for scanner.Scan() {
		currentString := scanner.Text()
		split := strings.Split(currentString, " ")
		var opponentsSelection selection
		if split[0] == "A" {
			opponentsSelection = Rock
		} else if split[0] == "B" {
			opponentsSelection = Paper
		} else {
			opponentsSelection = Scissors
		}
		myChoices := outcomeMap[opponentsSelection]
		var mySelection selection
		if split[1] == "X" {
			mySelection = Rock
		} else if split[1] == "Y" {
			mySelection = Paper
		} else {
			mySelection = Scissors
		}
		currentScore := int(mySelection) + myChoices[mySelection]
		totalScore += currentScore
	}
	log.Println(totalScore)
}

type selection int
const (
	Rock selection = 1
	Paper = 2
	Scissors = 3
)
