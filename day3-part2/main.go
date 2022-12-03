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
	count := 1
	totalLetterMap := make(map[string]int)
LineLoop:
	for scanner.Scan() {
		currentString := scanner.Text()
		currentLetterMap := make(map[string]int)
		for _, char := range(currentString)  {
			currentLetter := string(char)
			if _, ok := currentLetterMap[currentLetter]; !ok && count != 3 {
				currentLetterMap[currentLetter] = 1
				tlval, tlok := totalLetterMap[currentLetter]
				if tlok {
					totalLetterMap[currentLetter] = tlval + 1
				}else {
					totalLetterMap[currentLetter] = 1
				}
			}
			if count == 3 {
				if val, _ := totalLetterMap[currentLetter]; val == 2 {
					index := strings.Index(letterIndex, currentLetter) + 1
					sum += index
					totalLetterMap = make(map[string]int)
					count = 1
					continue LineLoop
				}
			}
		}
		count++
	}
	log.Println(sum)
}
