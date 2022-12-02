package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strconv"
	"sort"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	top3 := []int{-1, -1, -1}
	sort.Ints(top3)
	currentTotal := 0
	for scanner.Scan() {
		currentString := scanner.Text()
		if currentString == "" {
			if currentTotal > top3[0] {
				top3[0] = currentTotal
				sort.Ints(top3)
			}
			currentTotal = 0
		} else {
			current, _ := strconv.Atoi(currentString)
			currentTotal += current
		}

	}

	//catch last one
	if currentTotal > top3[0] {
		top3[0] = currentTotal
		sort.Ints(top3)
	}

	total := 0
	for _, top := range(top3) {
		total += top
	}
	fmt.Println(total)
}
