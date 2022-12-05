package main

import (
	"os"
	"bufio"
	"log"
	"strings"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		currentString := scanner.Text()
		commaSplit := strings.Split(currentString, ",")
		xSplit := strings.Split(commaSplit[0], "-")
		ySplit := strings.Split(commaSplit[1], "-")
		xMin , _ := strconv.Atoi(xSplit[0])
		xMax , _ := strconv.Atoi(xSplit[1])
		yMin , _ := strconv.Atoi(ySplit[0])
		yMax , _ := strconv.Atoi(ySplit[1])
		xLength := xMax - xMin
		yLength := yMax - yMin
		if xLength > yLength {
			if yMin >= xMin && yMax <= xMax {
				count++
			}
		}
		if yLength >= xLength {
			if xMin >= yMin && xMax <= yMax {
				count++
			}
		}
	}
	log.Println(count)
}
