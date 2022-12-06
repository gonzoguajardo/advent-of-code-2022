package main

import (
	"os"
	"bufio"
	"log"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var currentString string
	for scanner.Scan() {
		currentString = scanner.Text()
	}


	var hit int
Outer:
	for i := 0; i < len(currentString); i++ {
		currentMap := make(map[string]bool)
		currentLetter := currentString[i:i+1]
		for j := i ; j < i + 4 ; j++ {
			currentLetter = currentString[j:j+1]
			if _, ok := currentMap[currentLetter]; ok {
				continue Outer
			}
			currentMap[currentLetter] = true
		}
		hit = i
		break Outer
	}
	log.Println(hit + 4)
}
