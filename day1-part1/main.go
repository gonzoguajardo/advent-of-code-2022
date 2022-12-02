package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	max := -1
	currentTotal := 0
	for scanner.Scan() {
		currentString := scanner.Text()
		if currentString == "" {
			// fmt.Println(currentTotal)
			if currentTotal > max {
				max = currentTotal
			}
			currentTotal = 0
		} else {
			current, _ := strconv.Atoi(currentString)
			currentTotal += current
		}

	}
	fmt.Println(max)
}
