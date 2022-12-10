package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	matrix := make([][]int, 0)
	for scanner.Scan() {
		currentString := scanner.Text()
		currentRow := make([]int, 0)
		for i, _ := range(currentString) {
			currentNubmerString := currentString[i:i+1]
			currentNumber, _ := strconv.Atoi(currentNubmerString)
			currentRow = append(currentRow, currentNumber)
		}
		matrix = append(matrix, currentRow)
	}
	maxMatrix := make([][][4]int, len(matrix))
	for r := 0; r < len(matrix); r++ {
		maxMatrix[r] = make([][4]int, len(matrix[r]))
		for c := 0; c < len(matrix[0]); c++ {
			currentDirectionSlice := [4]int{-1, -1, -1, -1}
			if c - 1 >= 0 {
				currentNumber := matrix[r][c-1]
				currentLeftMax := maxMatrix[r][c-1][0]
				if currentNumber > currentLeftMax {
					currentLeftMax = currentNumber
				}
				currentDirectionSlice[0] = currentLeftMax
			}
			if r - 1 >= 0 {
				currentNumber := matrix[r-1][c]
				currentTopMax := maxMatrix[r-1][c][1]
				if currentNumber > currentTopMax {
					currentTopMax = currentNumber
				}
				currentDirectionSlice[1] = currentTopMax
			}
			maxMatrix[r][c] = currentDirectionSlice
		}

	}
	for r:= len(matrix) - 1; r >= 0; r-- {
		for c := len(matrix[0]) - 1; c >= 0; c-- {
			if c + 1 <= len(matrix[0]) - 1 {
				currentNumber := matrix[r][c+1]
				currentRightMax := maxMatrix[r][c+1][2]
				if currentNumber > currentRightMax {
					currentRightMax = currentNumber
				}
				maxMatrix[r][c][2]= currentRightMax
			}
			if r + 1 <= len(matrix) - 1{
				currentNumber := matrix[r+1][c]
				currentBottomMax := maxMatrix[r+1][c][3]
				if currentNumber > currentBottomMax {
					currentBottomMax = currentNumber
				}
				maxMatrix[r][c][3] = currentBottomMax
			}
		}
	}
	// [left, top, right, bottom]

	maxScenicScore := -1
	scenicScoreMatrix := make([][]int, 0)
	for r:= 1; r < len(matrix) - 1; r++ {
		rowScenicScoreMatrix := make([]int, 0)
		for c := 1; c < len(matrix[0]) - 1; c++ {
			currentNumber := matrix[r][c]
			currentScenicScore := 1
			var score int
			var scoreApplied bool

			// top
			score = 0
			scoreApplied = false
			for movingR := r - 1; movingR >=0; movingR-- {
				movingNumber := matrix[movingR][c]
				score++
				if movingNumber >= currentNumber {
					currentScenicScore *= score
					scoreApplied = true
					break
				}
			}
			if !scoreApplied {
				currentScenicScore *= score
			}

			// bottom
			score = 0
			scoreApplied = false
			for movingR := r + 1; movingR < len(matrix); movingR++ {
				movingNumber := matrix[movingR][c]
				score++
				if movingNumber >= currentNumber {
					currentScenicScore *= score
					scoreApplied = true
					break
				}
			}
			if !scoreApplied {
				currentScenicScore *= score
			}

			// Left
			score = 0
			scoreApplied = false
			for movingC := c - 1; movingC >=0; movingC-- {
				movingNumber := matrix[r][movingC]
				score++
				if movingNumber >= currentNumber {
					scoreApplied = true
					currentScenicScore *= score
					break
				}
			}
			if !scoreApplied {
				currentScenicScore *= score
			}

			// Right
			score = 0
			scoreApplied = false
			for movingC := c + 1; movingC < len(matrix[0]); movingC++ {
				movingNumber := matrix[r][movingC]
				score++
				if movingNumber >= currentNumber {
					scoreApplied = true
					currentScenicScore *= score
					break
				}
			}
			if !scoreApplied {
				currentScenicScore *= score
			}

			rowScenicScoreMatrix = append(rowScenicScoreMatrix, currentScenicScore)
			if currentScenicScore > maxScenicScore {
				maxScenicScore = currentScenicScore
			}
		}
		scenicScoreMatrix = append(scenicScoreMatrix, rowScenicScoreMatrix)
	}
	log.Println(maxScenicScore)
}

func print(matrix [][]int) {
	for _, row := range(matrix) {
		log.Println(row)
	}
}

func printMax(maxMatrix [][][4]int) {
	for _, row := range(maxMatrix) {
		log.Println(row)
	}
}
