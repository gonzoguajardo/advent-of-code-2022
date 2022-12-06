package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
	"strconv"
	"regexp"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var binsSetup, binsPopulated bool
	initialStackStrings := make([]string, 0)
	var bins [][]string

	for scanner.Scan() {
		currentString := scanner.Text()
		if !binsSetup && !binsPopulated {
			if !strings.Contains(currentString, "[") {
				binsSplit := strings.Split(currentString, " ")
				lastBinsSplit := binsSplit[len(binsSplit)-2]
				maxBins , _ := strconv.Atoi(lastBinsSplit)
				bins = make([][]string, maxBins)
				for n := 0; n < len(bins); n++ {
					bins[n] = make([]string, 0)
				}
				binsSetup = true
			} else {
				initialStackStrings = append(initialStackStrings, currentString)
			}
		} else if !binsPopulated {
			for lineIndex := len(initialStackStrings)-1; lineIndex >= 0; lineIndex -- {
				line := initialStackStrings[lineIndex]
				for n := 0; n < len(bins); n++ {
					currentIndex := 1 + n * 4
					letter := line[currentIndex:currentIndex+1]
					if letter != " " {
						bins[n] = append(bins[n], letter)
					}
				}
			}
			binsPopulated = true
		} else {
			r := regexp.MustCompile("[a-zA-Z ]+")
			numbers := r.Split(currentString, -1)
			itemsToMove, _ := strconv.Atoi(numbers[1])
			from, _ := strconv.Atoi(numbers[2])
			to, _ := strconv.Atoi(numbers[3])
			from--
			to--
			fromStartIndex := len(bins[from]) - itemsToMove
			itemsToAdd := bins[from][fromStartIndex:]
			for n := len(itemsToAdd) - 1; n >= 0; n-- {
				bins[to] = append(bins[to], itemsToAdd[n])
			}
			bins[from] = bins[from][:fromStartIndex]
		}
	}
	for _, bin := range(bins) {
		fmt.Print(bin[len(bin) - 1:][0])
	}
	fmt.Println()
}

