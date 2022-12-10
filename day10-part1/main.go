package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	defer file.Close()
	x := 1
	cycle := 1
	nextCycle := 20
	total := 0
	for scanner.Scan() {
		currentString := scanner.Text()
		isNoop := strings.Contains(currentString, "noop")
		if isNoop {
			cycle++
		} else {
			cycle++
			if cycle == nextCycle {
				total += cycle * x
				nextCycle += 40
			}
			addXSplit := strings.Split(currentString, "addx ")
			addNumber, _:= strconv.Atoi(addXSplit[1])
			x += addNumber
			cycle++
		}
		if cycle == nextCycle {
			total += cycle * x
			nextCycle += 40
		}
	}
	log.Println("total", total)
}
