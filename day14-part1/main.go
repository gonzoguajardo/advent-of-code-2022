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
	filledPointMap := createFilledPointMap()
	maxX := math.MinInt
	for k, _ := range(filledPointMap) {
		if k[1] > maxX {
			maxX = k[1]
		}
	}

	count := 0
Outer:
	for true {
		currentPosition := [2]int{500, 0}
		for true{
			if currentPosition[1] > maxX {
				break Outer
			}

			// move down
			newPoition := [2]int{currentPosition[0], currentPosition[1] + 1}
			if ok, _ := filledPointMap[newPoition]; !ok {
				currentPosition[1]++
				continue
			}

			// move down left
			newPoition = [2]int{currentPosition[0] - 1, currentPosition[1] + 1}
			if ok, _ := filledPointMap[newPoition]; !ok {
				currentPosition[0]--
				currentPosition[1]++
				continue
			}

			// move down right
			newPoition = [2]int{currentPosition[0] + 1, currentPosition[1] + 1}
			if ok, _ := filledPointMap[newPoition]; !ok {
				currentPosition[0]++
				currentPosition[1]++
				continue
			}

			filledPointMap[currentPosition] = true
			count++
			break
		}
	}
	log.Println(count)
}

func createFilledPointMap() map[[2]int]bool {
	filledPointMap := make(map[[2]int]bool)
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentString := scanner.Text()
		linePoints := strings.Split(currentString, " -> ")
		// log.Println(linePoints)
		currentLinePointSlice := make([][2]int, 0)
		currentLinePointSlice = append(currentLinePointSlice, makePointFromString(linePoints[0]))
		for index := 1 ; index < len(linePoints); index++ {
			endPoint := makePointFromString(linePoints[index])
			currentLinePointSlice = append(currentLinePointSlice, endPoint)
			startPoint := currentLinePointSlice[index-1]
			// log.Println(startPoint, endPoint)
			for true {
				if startPoint[0] != endPoint[0] {
					if startPoint[0] < endPoint[0] {
						filledPointMap[startPoint] = true
						startPoint[0]++
					} else {
						filledPointMap[endPoint] = true
						endPoint[0]++
					}
					continue
				}
				if startPoint[1] != endPoint[1] {
					if startPoint[1] < endPoint[1] {
						filledPointMap[startPoint] = true
						startPoint[1]++
					} else {
					filledPointMap[endPoint] = true
						endPoint[1]++
					}
					continue
				}
				break
			}
			filledPointMap[startPoint] = true
		}
	}
	return filledPointMap
}

func makePointFromString(pointString string) [2]int {
	pointSlice := make([]int, 0)
	valueString := strings.Split(pointString, ",")
	for _, vs := range(valueString) {
		value, _ := strconv.Atoi(vs)
		pointSlice = append(pointSlice, value)
	}
	pointArray := *(*[2]int)(pointSlice)
	return pointArray
}
