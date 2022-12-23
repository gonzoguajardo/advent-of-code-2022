package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	s := createMatrix()
	log.Println("initial state")
	s.print()
	n := 1
	for i := 0; i < n; i++ {
		s.propose()
		s.move()
		log.Println("after round", i + 1)
		s.print()
	}
}

func (s *Solution) shiftDirections() {
	shift := s.currentDirection[1:]
	shift = append(shift, s.currentDirection[0])
	shiftArray := *(*[4]Direction)(shift)
	s.currentDirection = shiftArray
}

func (s *Solution) move() {
	for k, v := range(s.proposedMoves) {
		if len(v) == 1 {
			s.matrix[k[0]][k[1]] = "#"
			s.matrix[v[0][0]][v[0][1]] = "."
		}
	}
	s.shiftDirections()
}

func (s *Solution) propose() {
	s.proposedMoves = make(map[[2]int][][2]int)
	for r := range(s.matrix) {
		for c := range(s.matrix[r]) {
			if s.matrix[r][c] == "#" {
				currentPoint := [2]int{r, c}
				// log.Println("found", r, c)
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
	}
}

func (s *Solution) isThereSurrouningElf(check [2]int) bool {
	r := check[0]
	c := check[1]
	if r != 0 {
		if s.matrix[r-1][c] == "#" {return true}
		if c != 0 && s.matrix[r-1][c-1] == "#" {return true}
		if c != len(s.matrix[0]) - 1 && s.matrix[r-1][c+1] == "#" {return true}
	}
	if r != len(s.matrix) - 1 {
		if s.matrix[r+1][c] == "#" {return true}
		if c != 0 && s.matrix[r+1][c-1] == "#" {return true}
		if c != len(s.matrix[0]) - 1 && s.matrix[r+1][c+1] == "#" {return true}
	}
	if c != 0 {
		if s.matrix[r][c-1] == "#" {return true}
		if r != 0 && s.matrix[r-1][c-1] == "#" {return true}
		if r != len(s.matrix) - 1 && s.matrix[r+1][c-1] == "#" {return true}
	}
	if c != len(s.matrix[0]) - 1 {
		if s.matrix[r][c+1] == "#" {return true}
		if r != 0 && s.matrix[r-1][c+1] == "#" {return true}
		if r != len(s.matrix) - 1 && s.matrix[r+1][c+1] == "#" {return true}
	}
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
	if r == 0 || c == 0 || c == len(s.matrix[0]) - 1{
		return false
	}
	northEast := s.matrix[r-1][c-1]
	north := s.matrix[r-1][c]
	northWest := s.matrix[r-1][c+1]
	return northEast != "#" && north != "#" && northWest != "#"
}

func (s *Solution) canMoveSouth(check [2]int) bool {
        r := check[0]
	c := check[1]
	if r == len(s.matrix) - 1 || c == 0 || c == len(s.matrix[0]) - 1{
		return false
	}
	southEast := s.matrix[r+1][c-1]
	south := s.matrix[r+1][c]
	southWest := s.matrix[r+1][c+1]
	return southEast != "#" && south != "#" && southWest != "#"
}

func (s *Solution) canMoveWest(check [2]int) bool {
	r := check[0]
	c := check[1]
	if c == 0 || r == 0 || r == len(s.matrix) - 1 {
		return false
	}
	westNorth := s.matrix[r-1][c-1]
	west := s.matrix[r][c-1]
	westSouth := s.matrix[r+1][c-1]
	return westNorth != "#" && west != "#" && westSouth != "#"
}

func (s *Solution) canMoveEast(check [2]int) bool {
	r := check[0]
	c := check[1]
	if c == len(s.matrix[0]) - 1 || r == 0 || r == len(s.matrix) - 1{
		return false
	}
	eastNorth := s.matrix[r-1][c+1]
	east := s.matrix[r][c+1]
	eastSouth := s.matrix[r+1][c+1]
	return eastNorth != "#" && east != "#" && eastSouth != "#"
}

func (s *Solution) print() {
	for _, row := range(s.matrix) {
		log.Println(row)
	}
	log.Println()
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
	matrix := make([][]string, 0)
	for _, input := range(inputs) {
		current := make([]string, 0)
		for i := range(input) {
			current = append(current, input[i:i+1])
		}
		matrix = append(matrix, current)
	}
	return &Solution{
		matrix: matrix,
		currentDirection: [4]Direction{North, South, West, East},
		proposedMoves: make(map[[2]int][][2]int),
	}
}

type Solution struct {
	matrix [][]string
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
