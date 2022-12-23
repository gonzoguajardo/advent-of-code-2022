package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"fmt"
)

func main() {
    s := createMatrix()
	// log.Println("initial state")
	// s.print()
	n := 0
	for true {
		s.propose()
		n++
		if !s.move() {
			break
		}
		// log.Println("after round", i + 1)
		// s.print()
	}
	// log.Println(s.countEmptyGroundTilesInnerRectangle())
	log.Println(n)
	// s.print()
}

func (s *Solution) shiftDirections() {
	shift := s.currentDirection[1:]
	shift = append(shift, s.currentDirection[0])
	shiftArray := *(*[4]Direction)(shift)
	s.currentDirection = shiftArray
}

func (s *Solution) move() bool {
	moved := false
	for k, v := range(s.proposedMoves) {
		if len(v) == 1 {
			delete(s.currentPositions, v[0])
			s.currentPositions[k] = true
			moved = true
		}
	}
	s.shiftDirections()
	return moved
}

func (s *Solution) propose() {
	s.proposedMoves = make(map[[2]int][][2]int)
	for currentPoint, _ := range(s.currentPositions) {
	DirectionLoop:
		for _, d := range(s.currentDirection) {
			// log.Println("checking direction", d)
			if !s.isThereSurrouningElf(currentPoint) {continue}
			switch(d){
				case North:
				if s.canMoveNorth(currentPoint) {
					// log.Println("can move north")
					s.addProposedMove(currentPoint, d)
					break DirectionLoop
				}
				case South:
				if s.canMoveSouth(currentPoint) {
					// log.Println("can move south")
					s.addProposedMove(currentPoint, d)
					break DirectionLoop
				}
				case West:
				if s.canMoveWest(currentPoint) {
					// log.Println("can move west")
					s.addProposedMove(currentPoint, d)
					break DirectionLoop
				}
				case East:
				if s.canMoveEast(currentPoint) {
					// log.Println("can move east")
					s.addProposedMove(currentPoint, d)
					break DirectionLoop
				}
				default:
				log.Println("not implemented in propose using direction", d)
			}
		}
	}
}

func (s *Solution) isThereSurrouningElf(check [2]int) bool {
	r := check[0]
	c := check[1]
	if _, ok := s.currentPositions[[2]int{r-1, c}]; ok {return true}
	if _, ok := s.currentPositions[[2]int{r-1, c-1}]; ok {return true}
	if _, ok := s.currentPositions[[2]int{r-1, c+1}]; ok {return true}

	if _, ok := s.currentPositions[[2]int{r+1, c}]; ok {return true}
	if _, ok := s.currentPositions[[2]int{r+1, c-1}]; ok {return true}
	if _, ok := s.currentPositions[[2]int{r+1, c+1}]; ok {return true}

	if _, ok := s.currentPositions[[2]int{r, c-1}]; ok {return true}
	if _, ok := s.currentPositions[[2]int{r-1, c-1}]; ok {return true}
	if _, ok := s.currentPositions[[2]int{r+1, c-1}]; ok {return true}

	if _, ok := s.currentPositions[[2]int{r, c+1}]; ok {return true}
	if _, ok := s.currentPositions[[2]int{r-1, c+1}]; ok {return true}
	if _, ok := s.currentPositions[[2]int{r+1, c+1}]; ok {return true}

	return false
}

func (s *Solution) addProposedMove(currentPoint [2]int, d Direction) {
	r := currentPoint[0]
	c := currentPoint[1]
	var proposedMove [2]int
	switch(d){
		case North:
		proposedMove = [2]int{r-1, c}
		case South:
		proposedMove = [2]int{r+1, c}
		case West:
		proposedMove = [2]int{r, c-1}
		case East:
		proposedMove = [2]int{r, c+1}
		default:
		log.Println("not implemented in add proposed move")
	}
	s.proposedMoves[proposedMove] = append(s.proposedMoves[proposedMove], currentPoint)
}

func (s *Solution) canMoveNorth(check [2]int) bool {
	r := check[0]
	c := check[1]
	if _, ok := s.currentPositions[[2]int{r-1, c}]; ok {return false}
	if _, ok := s.currentPositions[[2]int{r-1, c-1}]; ok {return false}
	if _, ok := s.currentPositions[[2]int{r-1, c+1}]; ok {return false}
	return true
}

func (s *Solution) canMoveSouth(check [2]int) bool {
	r := check[0]
	c := check[1]
	if _, ok := s.currentPositions[[2]int{r+1, c}]; ok {return false}
	if _, ok := s.currentPositions[[2]int{r+1, c-1}]; ok {return false}
	if _, ok := s.currentPositions[[2]int{r+1, c+1}]; ok {return false}
	return true
}

func (s *Solution) canMoveWest(check [2]int) bool {
	r := check[0]
	c := check[1]
	if _, ok := s.currentPositions[[2]int{r, c-1}]; ok {return false}
	if _, ok := s.currentPositions[[2]int{r-1, c-1}]; ok {return false}
	if _, ok := s.currentPositions[[2]int{r+1, c-1}]; ok {return false}
	return true
}

func (s *Solution) canMoveEast(check [2]int) bool {
	r := check[0]
	c := check[1]
	if _, ok := s.currentPositions[[2]int{r, c+1}]; ok {return false}
	if _, ok := s.currentPositions[[2]int{r-1, c+1}]; ok {return false}
	if _, ok := s.currentPositions[[2]int{r+1, c+1}]; ok {return false}
	return true
}

func (s *Solution) countEmptyGroundTilesInnerRectangle() int {
	minR := math.MaxInt
	minC := math.MaxInt
	maxR := math.MinInt
	maxC := math.MinInt
	for k, _ := range(s.currentPositions) {
		r := k[0]
		c := k[1]
		if r < minR {
			minR = r
		}
		if r > maxR {
			maxR = r
		}
		if c < minC {
			minC = c
		}
		if c > maxC {
			maxC = c
		}
	}
	count := 0
	for r := minR; r <= maxR; r++ {
		for c := minC; c <= maxC; c++ {
			if _, ok := s.currentPositions[[2]int{r, c}]; ok {
			} else {
				count++
			}
		}
	}
	return count
}

func (s *Solution) print() {
	minR := math.MaxInt
	minC := math.MaxInt
	maxR := math.MinInt
	maxC := math.MinInt
	for k, _ := range(s.currentPositions) {
		r := k[0]
		c := k[1]
		if r < minR {
			minR = r
		}
		if r > maxR {
			maxR = r
		}
		if c < minC {
			minC = c
		}
		if c > maxC {
			maxC = c
		}
	}
	// log.Println("printing with", minR, maxR, minC, maxC)
	for r := minR; r <= maxR; r++ {
		for c := minC; c <= maxC; c++ {
			if _, ok := s.currentPositions[[2]int{r, c}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func createMatrix() *Solution {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	inputs := make([]string, 0)
	for scanner.Scan() {
		currentString := scanner.Text()
		inputs = append(inputs, currentString)
	}
	currentPositions := make(map[[2]int]bool)
	r := 0
	for _, input := range(inputs) {
		for i := range(input) {
			if input[i:i+1] == "#" {
				currentPositions[[2]int{r, i}] = true
			}
		}
		r++
	}
	return &Solution{
		currentPositions: currentPositions,
		currentDirection: [4]Direction{North, South, West, East},
		proposedMoves: make(map[[2]int][][2]int),
	}
}

type Solution struct {
	currentPositions map[[2]int]bool
	proposedMoves map[[2]int][][2]int
	currentDirection [4]Direction
}

type Direction int

const (
	North Direction = 0
	South Direction = 1
	West Direction = 2
	East Direction = 3
)
