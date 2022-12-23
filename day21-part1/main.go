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
	solvedMonkeys := make(map[string]int)
	problemMonkeys := make(map[string][]string)
	for scanner.Scan() {
		currentString := scanner.Text()
		colonSplit := strings.Split(currentString, ": ")
		// log.Println(colonSplit[0])
		spaceSplit := strings.Split(colonSplit[1], " ")
		if len(spaceSplit) == 1 {
			value, _ := strconv.Atoi(spaceSplit[0])
			solvedMonkeys[colonSplit[0]] = value
		} else {
			problemMonkeys[colonSplit[0]] = spaceSplit
		}
	}

	for len(problemMonkeys) > 0 {
		for k, v := range(problemMonkeys) {
			leftV, leftOk := solvedMonkeys[v[0]]
			rightV, rightOk := solvedMonkeys[v[2]]
			if leftOk && rightOk {
				operation := Operation(v[1])
				var answer int
				switch(operation){
					case Addition:
					answer = leftV + rightV
					case Subtraction:
					answer = leftV - rightV
					case Multiply:
					answer = leftV * rightV
					case Divide:
					answer = leftV / rightV
					default:
					log.Println("operation not impemented")
				}
				solvedMonkeys[k] = answer
				delete(problemMonkeys, k)
			}
		}
	}
	log.Println(solvedMonkeys["root"])
}

type Operation string

const (
	Addition Operation = "+"
	Subtraction Operation = "-"
	Multiply Operation = "*"
	Divide Operation = "/"
)
