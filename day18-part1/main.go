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
	defer file.Close()
	scanner := bufio.NewScanner(file)
	cubes := make([][3]int, 0)
	for scanner.Scan() {
		currentString := scanner.Text()
		currentCubeSlice := make([]int, 0)
		commaSplit := strings.Split(currentString, ",")
		for _, s := range(commaSplit) {
			value, _ := strconv.Atoi(s)
			currentCubeSlice = append(currentCubeSlice, value)
		}
		currentCubeArray := *(*[3]int)(currentCubeSlice)
		cubes = append(cubes, currentCubeArray)
	}
	s := Solution(cubes)
	log.Println(s)
}

func Solution(cubes [][3]int) int {
	surfaceArea := 0
	cubeSet := make(map[[3]int]bool)
	for _, cube := range(cubes) {
		cubeSet[cube] = true
	}
	for k, _ := range(cubeSet) {
		for _, side := range(GetAllSides(k)) {
			if ok, _ := cubeSet[side]; !ok {
				surfaceArea++
			}
		}
	}
	return surfaceArea
}

func GetAllSides(cube [3]int) [][3]int{
	sides := make([][3]int, 0)
	for axis := range(cube) {
		currentSidePlus := cube
		currentSideMinus := cube
		currentSideMinus[axis]--
		currentSidePlus[axis]++
		sides = append(sides, currentSideMinus)
		sides = append(sides, currentSidePlus)
	}
	return sides
}
