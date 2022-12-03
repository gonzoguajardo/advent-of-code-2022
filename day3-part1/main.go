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
	letterIndex := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	sum := 0
LineLoop:
	for scanner.Scan() {
		currentString := scanner.Text()
		currentLetterMap := make(map[string]int)
		for x, char := range(currentString)  {
			currentLetter := string(char)
			if x < len(currentString) / 2 {
				currentLetterMap[currentLetter] = 1
			} else {
				if _, ok := currentLetterMap[currentLetter]; ok {
					index := strings.Index(letterIndex, currentLetter) + 1
					sum += index
					continue LineLoop
				}
			}
		}
	}
	log.Println(sum)
}
