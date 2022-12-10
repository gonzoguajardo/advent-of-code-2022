package main

import (
	"bufio"
	"log"
	"math"
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

	stringMatrix := make([][]string, 0)
	stringMatrix = append(stringMatrix, make([]string, 0))
	klass := &Klass{
		x: 1,
		cycle: 1,
		nextCycle: 41,
		total: 0,
		currentRow: 0,
		stringMatrix: stringMatrix,
	}

	for scanner.Scan() {
		currentString := scanner.Text()
		isNoop := strings.Contains(currentString, "noop")
		if isNoop {
			klass.applyCycle()
		} else {
			klass.applyCycle()
			addXSplit := strings.Split(currentString, "addx ")
			addNumber, _:= strconv.Atoi(addXSplit[1])
			klass.applyCycle()
			klass.addX(addNumber)
		}
	}
	for _, row := range(klass.stringMatrix) {
		log.Println(row)
	}
}

func (klass *Klass) applyCycle() {
	if math.Abs(float64(klass.x - (klass.cycle % 40 - 1))) <=1 {
		klass.stringMatrix[klass.currentRow] = append(klass.stringMatrix[klass.currentRow], "#")
	} else {
		klass.stringMatrix[klass.currentRow] = append(klass.stringMatrix[klass.currentRow], ".")
	}
	klass.cycle++
	if klass.cycle == klass.nextCycle {
		klass.total += klass.cycle * klass.x
		klass.nextCycle += 40
		klass.currentRow++
		klass.stringMatrix = append(klass.stringMatrix, make([]string, 0))
	}
}

func (klass *Klass) addX(addX int) {
	klass.x += addX
}

type Klass struct {
	x int
	cycle int
	nextCycle int
	total int
	currentRow int
	stringMatrix [][]string
}
