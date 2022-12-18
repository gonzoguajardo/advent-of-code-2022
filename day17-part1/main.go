package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("rocks.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	rocks := make([][]string, 0)
	currentRock := make([]string, 0)
	for scanner.Scan() {
		currentString := scanner.Text()
		if currentString == "" {
			rocks = append(rocks, currentRock)
			currentRock = make([]string, 0)
		}else {
			currentRock = append(currentRock, currentString)
		}
	}
	rocks = append(rocks, currentRock)
	matrix := newMatrix()




	file, err = os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner = bufio.NewScanner(file)
	var directionString string
	for scanner.Scan() {
		directionString = scanner.Text()
	}
	currentRockIndex := 0
	doneFallingRocks := true
	addedRockCount := 0
Outer:
	for true {
		for x := range(directionString) {
			if doneFallingRocks {
				addedRockCount++
				if addedRockCount == 2023 {
					break Outer
				}
				matrix.AddRock(rocks[currentRockIndex])
				currentRockIndex++
				if currentRockIndex >= len(rocks) {
					currentRockIndex = 0
				}
				doneFallingRocks = false
				// log.Println("adding rocks", addedRockCount)
				// matrix.Print()
			}
			char := directionString[x:x+1]
			currentDir := Direction(char)
			// log.Println("moving current dir", currentDir)
			matrix.moveRocks(currentDir)
			// matrix.Print()
			doneFallingRocks = matrix.fallRocks()
		}
	}
	// matrix.Print()
	log.Println(matrix.maxHeight)
	log.Println(len(matrix.Chamber) - 4)
	log.Println("rocks added", addedRockCount)
}

func (matrix *matrix) canMoveRocks(dir Direction) bool {
	for hIndex := matrix.currentHeight; hIndex < len(matrix.Chamber); hIndex++ {
		found := false
		h := matrix.Chamber[hIndex]
		var copyH [7]string
		copy(copyH[:], h[:])
		switch dir {
		case Right:
			for x := len(h)-2; x >= 0; x-- {
				if h[x] == "@" {
					found = true
					if copyH[x+1] == "." {
						copyH[x+1] = "@"
						copyH[x] = "."
					}else {
						return false
					}
				}
			}
		case Left:
			for x := 1; x < len(h); x++ {
				if h[x] == "@" {
					found = true
					if copyH[x-1] == "." {
						copyH[x-1] = "@"
						copyH[x] = "."
					} else {
						return false
					}
				}
			}
		}
		if !found {
			return true
		}
	}
	return true
}

func (matrix *matrix) moveRocks(dir Direction) {
	if !matrix.canMoveRocks(dir) {
		return
	}
	for hIndex , h := range(matrix.Chamber) {
		var copyH [7]string
		copy(copyH[:], h[:])
		switch dir {
		case Right:
			for x := len(h)-2; x >= 0; x-- {
				if h[x] == "@" {
					if copyH[x+1] == "." {
						copyH[x+1] = "@"
						copyH[x] = "."
					}else {
						return
					}
				}
			}
		case Left:
			for x := 1; x < len(h); x++ {
				if h[x] == "@" {
					if copyH[x-1] == "." {
						copyH[x-1] = "@"
						copyH[x] = "."
					} else {
						return
					}
				}
			}
		}
		matrix.Chamber[hIndex] = copyH
	}
}

func (matrix *matrix) stopRocks() {
	// log.Println("stopping rocks", matrix.currentHeight, matrix.maxHeight)
	// matrix.Print()
	nextHeight := 0
	for hIndex := matrix.currentHeight; hIndex < len(matrix.Chamber); hIndex++ {
		currentH := matrix.Chamber[hIndex]
		foundRock := false
		for x := 0; x < len(currentH); x++ {
			if currentH[x] == "@" {
				matrix.Chamber[hIndex][x] = "#"
				foundRock = true
			}
		}
		if foundRock && hIndex >= matrix.maxHeight {
			nextHeight++
		}
	}
	matrix.maxHeight = matrix.maxHeight + nextHeight
	// log.Println("after stopping rocks", matrix.currentHeight, matrix.maxHeight)
}

func (matrix *matrix) canStopRocks() bool {
	if matrix.currentHeight == 0 {
		return true
	} else {
		hIndex := matrix.currentHeight
		for {
			currentH := matrix.Chamber[hIndex]
			// log.Println("inside can stop rock", hIndex, currentH)
			found := false
			for x := 0; x < len(currentH); x++ {
				if currentH[x] == "@" {
					found = true
					if matrix.Chamber[hIndex-1][x] == "#" {
						return true
					}
				}
			}
			if !found {
				return false
			}
			hIndex++
		}
	}
}

func (matrix *matrix) fallRocks() bool {
	if matrix.currentHeight <= matrix.maxHeight {
		stopRocks := matrix.canStopRocks()
		// log.Println("checking stop rocks", stopRocks, matrix.currentHeight)
		// matrix.Print()
		if stopRocks {
			// log.Println("stopping rocks")
			// matrix.Print()
			matrix.stopRocks()
			return true
		}
	}
	for hIndex := matrix.currentHeight; hIndex < len(matrix.Chamber); hIndex++ {
		currentH := matrix.Chamber[hIndex]
		for x := 0; x < len(currentH); x++ {
			if currentH[x] == "@" {
				if matrix.Chamber[hIndex-1][x] == "." {
					matrix.Chamber[hIndex][x] = "."
					matrix.Chamber[hIndex-1][x] = "@"
				}
			}
		}

	}
	matrix.currentHeight--
	return false
}

type matrix struct {
	size int
	currentHeight int
	maxHeight int
	Chamber [][7]string
}

func (matrix matrix) Print() {
	for h := len(matrix.Chamber) - 1; h >= 0; h-- {
		log.Println(matrix.Chamber[h])
	}
}

func newMatrix() *matrix {
	chamber := make([][7]string, 3)
	for x := range(chamber) {
		chamber[x] = [7]string{".", ".", ".", ".", ".", ".", "."}
	}
	return &matrix{
		size: 3,
		currentHeight: 0,
		maxHeight: 0,
		Chamber: chamber,
	}
}

func (matrix *matrix) AddRock(rock []string) {
	// log.Println("adding rock", rock)
	currentHIndex := matrix.maxHeight + 2
	matrix.currentHeight = currentHIndex + 1
	for currentSize := len(matrix.Chamber); currentSize <= currentHIndex+ len(rock) ; currentSize++ {
		matrix.Chamber = append(matrix.Chamber, [7]string{".", ".", ".", ".", ".", ".", "."})
	}
	for rockIndex := len(rock) - 1; rockIndex >= 0; rockIndex -- {
		rockH := rock[rockIndex]
		currentHIndex++
		newH := make([]string, 7)
		newH[0] = "."
		newH[1] = "."
		x := 1
		for rockX := range(rockH) {
			x++
			currentChar := rockH[rockX:rockX+1]
			if currentChar == "#" {
				newH[x] = "@"
			}else {
				newH[x] = currentChar
			}
		}
		for x <= 5 {
			x++
			newH[x] = "."
		}
		newHSlice := *(*[7]string)(newH)
		matrix.Chamber[currentHIndex] = newHSlice
	}
}

type Direction string

const (
	Right Direction = ">"
	Left Direction = "<"
)
